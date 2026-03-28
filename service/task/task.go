package task

import (
	"context"
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	stlmaps "github.com/kkkunny/stl/container/maps"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/container/tuple"
	stlerr "github.com/kkkunny/stl/error"
	stlos "github.com/kkkunny/stl/os"
	stlval "github.com/kkkunny/stl/value"
	xldto "github.com/kkkunny/xunlei/dto"
	"gorm.io/gorm"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/dal/db"
	"github.com/kkkunny/MDM/dal/db/po"
	"github.com/kkkunny/MDM/dal/qb"
	"github.com/kkkunny/MDM/dal/xl"
	"github.com/kkkunny/MDM/model/dto"
	"github.com/kkkunny/MDM/model/vo"
	"github.com/kkkunny/MDM/service/xltorrent"
	"github.com/kkkunny/MDM/util"
)

// GetAllTasks 获取所有任务
func GetAllTasks(ctx context.Context, forceUpdate ...bool) ([]dto.Task, error) {
	if stlslices.Last(forceUpdate) {
		return tasksCache.GetLatest(ctx)
	}
	return tasksCache.Get(ctx)
}

// GetTaskByHash 根据hash获取任务
func GetTaskByHash(ctx context.Context, hash string, forceUpdate ...bool) (dto.Task, error) {
	tasks, err := GetAllTasks(ctx, forceUpdate...)
	if err != nil {
		return nil, err
	}
	task, ok := stlslices.FindFirst(tasks, func(i int, t dto.Task) bool {
		return t.Hash() == hash
	})
	if !ok {
		return nil, stlerr.Errorf("not found task by hash `%s`", hash)
	}
	return task, nil
}

var taskAutoManagerLocker sync.Mutex

// AutoManageTasks 自动管理任务
func AutoManageTasks(ctx context.Context) error {
	if !taskAutoManagerLocker.TryLock() {
		return nil
	}
	defer taskAutoManagerLocker.Unlock()

	tasks, err := GetAllTasks(ctx)
	if err != nil {
		return err
	}

	// 下载完成的任务
	completedTasks := stlslices.Map(stlslices.Filter(tasks, func(_ int, t dto.Task) bool {
		xlTask, ok := t.(*dto.XLTask)
		if !ok {
			return false
		}
		if xlTask.TaskInfo.Phase != xldto.TaskPhaseTypeComplete {
			return false
		}

		// 跳过保存地址不存在的任务，有可能上次迁移时失败了
		exist, err := stlerr.ErrorWith(stlos.Exist(xlTask.SavePath()))
		if err != nil {
			_ = config.Logger.Warn(err)
			return false
		}
		if exist {
			// 跳过保存地址里存在迅雷临时文件的任务
			var existTempFile bool
			err = stlerr.ErrorWrap(filepath.WalkDir(xlTask.SavePath(), func(_ string, entry fs.DirEntry, err error) error {
				if err != nil {
					return stlerr.ErrorWrap(err)
				}
				if entry.IsDir() {
					return nil
				}
				existTempFile = existTempFile || strings.HasSuffix(strings.ToLower(entry.Name()), ".xltd")
				return nil
			}))
			if err != nil {
				_ = config.Logger.Warn(err)
				return false
			}
			if existTempFile {
				return false
			}
		}

		return true
	}), func(_ int, t dto.Task) *dto.XLTask { return t.(*dto.XLTask) })
	if len(completedTasks) > 0 {
		return Completed(ctx, completedTasks...)
	}

	// 下载阻塞的任务
	cloggedTasks := stlslices.Map(stlslices.Filter(tasks, func(_ int, t dto.Task) bool {
		xlTask, ok := t.(*dto.XLTask)
		if !ok {
			return false
		}
		if xlTask.TaskInfo.Phase == xldto.TaskPhaseTypeComplete || xlTask.TaskInfo.Progress > 99 {
			return false
		} else if xlTask.TaskInfo.Phase == xldto.TaskPhaseTypeError {
			return true
		}
		dur, speed := time.Since(xlTask.TaskInfo.CreatedTime), stlos.Size(xlTask.TaskInfo.Speed)*stlos.Byte
		switch {
		case dur > time.Minute && speed < stlos.KiB:
			// 超过一分钟，速度小于1KB的
			return true
		}
		return false
	}), func(_ int, t dto.Task) *dto.XLTask { return t.(*dto.XLTask) })
	if len(cloggedTasks) > 0 {
		return Clogged(ctx, cloggedTasks...)
	}

	return nil
}

// Completed 下载完成
func Completed(ctx context.Context, tasks ...*dto.XLTask) error {
	_ = config.Logger.Infof("tasks `%v` download completed", stlslices.Map(tasks, func(_ int, t *dto.XLTask) string { return t.Name() }))

	// 查找种子文件
	hash2TaskAndTorrentPath := stlslices.ToMap(tasks, func(t *dto.XLTask) (string, tuple.Tuple2[*dto.XLTask, string]) {
		return t.Hash(), tuple.Pack2[*dto.XLTask, string](t, "")
	})
	torrentMIs := xltorrent.TorrentsCache.Get()
	for fp, tmi := range torrentMIs {
		hash := tmi.HashInfoBytes().HexString()
		tt, ok := hash2TaskAndTorrentPath[hash]
		if !ok {
			continue
		}
		hash2TaskAndTorrentPath[hash] = tuple.Pack2(tt.E1(), fp)
	}

	// 迁移
	for _, taskAndTorrentPath := range hash2TaskAndTorrentPath {
		task, tp := taskAndTorrentPath.Unpack()
		if tp == "" {
			_ = config.Logger.Warnf("can not find torrent for task `%s`", task.Name())
		}

		// 迁移文件
		resources, err := stlerr.ErrorWith(xl.Client.ListResource(ctx, task.URL))
		if err != nil {
			return err
		}
		singleFile := !stlslices.First(resources).IsDir()
		filename := stlslices.First(resources).GetName()
		fromPath := task.SavePath()
		if singleFile {
			// 单文件时迅雷会自动创建一个文件夹
			fromPath = filepath.Join(fromPath, filename)
		}
		toPath := filepath.Join(config.DownloadCompleteDir, filename)
		exist, err := stlerr.ErrorWith(stlos.Exist(task.SavePath()))
		if err != nil {
			return err
		}
		if exist {
			err = stlerr.ErrorWrap(os.Rename(fromPath, toPath))
			if err != nil {
				_ = config.Logger.Warnf("fromPath=%s, toPath=%s, xltask=%s", fromPath, toPath, util.ToJson[string](task.TaskInfo))
				return err
			}
			// 单文件时迅雷会自动创建一个文件夹，需要删除
			if singleFile {
				err = stlerr.ErrorWrap(os.RemoveAll(filepath.Dir(fromPath)))
				if err != nil {
					return err
				}
			}
		}

		// 新建qb任务
		var existedQbTask *dto.QBTask
		if existedTask := stlval.IgnoreWith(GetTaskByHash(ctx, task.Hash())); existedTask != nil {
			existedQbTask, _ = existedTask.(*dto.QBTask)
		}
		if existedQbTask == nil {
			if tp != "" {
				err = stlerr.ErrorWrap(qb.Client.AddTorrentFromFileCtx(ctx, tp, map[string]string{
					"skip_checking": "true",
					"rename":        task.Name(),
				}))
			} else {
				err = stlerr.ErrorWrap(qb.Client.AddTorrentFromUrlCtx(ctx, task.URL, map[string]string{
					"skip_checking": "true",
					"rename":        task.Name(),
				}))
			}
			if err != nil {
				return err
			}
		}

		// 删除数据库
		d, err := db.NewTasksDal(ctx)
		if err != nil {
			return err
		}
		if err = d.DelByIDs(task.ID()); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// 删除迅雷任务
		err = stlerr.ErrorWrap(xl.Client.DeleteTask(ctx, task.TaskInfo.ID, false))
		if err != nil {
			return err
		}

		// 回调
		if config.TaskDownloadCompletedFallbackAddr != "" {
			fallbackReq := &vo.TaskDownloadCompletedFallbackRequest{
				Name:     task.Name(),
				Hash:     task.Hash(),
				SavePath: toPath,
			}
			httpReq, err := stlerr.ErrorWith(http.NewRequestWithContext(ctx, http.MethodPost, config.TaskDownloadCompletedFallbackAddr, strings.NewReader(util.ToJson[string](fallbackReq))))
			if err != nil {
				_ = config.Logger.Warn(err)
			} else {
				httpReq.Header.Set("Content-Type", "application/json")
				resp, err := stlerr.ErrorWith(http.DefaultClient.Do(httpReq))
				if err != nil {
					_ = config.Logger.Warn(err)
				} else {
					defer resp.Body.Close()
					if resp.StatusCode != http.StatusOK {
						_ = config.Logger.Warnf("fallback request failed, statusCode=%d", resp.StatusCode)
					}
				}
			}
		}
	}
	return nil
}

// Clogged 下载阻塞
func Clogged(ctx context.Context, tasks ...*dto.XLTask) error {
	_ = config.Logger.Tracef("tasks `%v` download clogged", stlslices.Map(tasks, func(_ int, t *dto.XLTask) string { return t.Name() }))

	d, err := db.NewTasksDal(ctx)
	if err != nil {
		return err
	}
	dbTasks, err := d.MGetByIDs(stlslices.Map(tasks, func(_ int, t *dto.XLTask) string { return t.ID() })...)
	if err != nil {
		return err
	}
	dbTasks = stlmaps.Filter(dbTasks, func(k string, v *po.Task) bool {
		return strings.Count(*v.AvailableLinks, ",") > 0
	})
	tasks = stlslices.Filter(tasks, func(_ int, t *dto.XLTask) bool {
		return stlmaps.ContainKey(dbTasks, t.ID())
	})

	if len(tasks) == 0 {
		return nil
	}
	_ = config.Logger.Infof("tasks `%v` download clogged, need using backup link", stlslices.Map(tasks, func(_ int, t *dto.XLTask) string { return t.Name() }))

	for _, task := range tasks {
		// 删除迅雷任务
		err = stlerr.ErrorWrap(xl.Client.DeleteTask(ctx, task.TaskInfo.ID, false))
		if err != nil {
			return err
		}
		err = stlerr.ErrorWrap(os.RemoveAll(task.SavePath()))
		if err != nil {
			return err
		}

		// 新建迅雷任务
		dbTask := dbTasks[task.ID()]
		availableLinks := util.FromJson[string, []string](*dbTask.AvailableLinks)
		newTask, err := stlerr.ErrorWith(xl.Client.CreateTask(ctx, task.TaskInfo.Name, availableLinks[1]))
		if err != nil {
			return err
		}

		// 更新数据库
		dbTask.Xlid = &newTask.ID
		dbTask.AvailableLinks = stlval.Ptr(util.ToJson[string](availableLinks[1:]))
		unavailableLinks := util.FromJson[string, []string](*dbTask.UnavailableLinks)
		dbTask.UnavailableLinks = stlval.Ptr(util.ToJson[string](append(unavailableLinks, availableLinks[0])))
		if err = d.MSave(dbTask); err != nil {
			return err
		}
	}
	return nil
}
