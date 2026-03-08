package task

import (
	"context"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/anacrolix/torrent/metainfo"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/container/tuple"
	stlerr "github.com/kkkunny/stl/error"
	stlos "github.com/kkkunny/stl/os"
	stlval "github.com/kkkunny/stl/value"
	xldto "github.com/kkkunny/xunlei/dto"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/dal/qb"
	"github.com/kkkunny/MDM/dal/xl"
	"github.com/kkkunny/MDM/model/dto"
	"github.com/kkkunny/MDM/model/vo"
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

	tasks, err := GetAllTasks(ctx, true)
	if err != nil {
		return err
	}

	// 过滤完成的任务
	downTasks := stlslices.FlatMap(tasks, func(_ int, t dto.Task) []*dto.XLTask {
		xlTask, ok := t.(*dto.XLTask)
		if !ok {
			return nil
		}
		if xlTask.TaskInfo.Phase != xldto.TaskPhaseTypeComplete {
			return nil
		}

		// 跳过保存地址不存在的任务，有可能上次迁移时失败了
		exist, err := stlerr.ErrorWith(stlos.Exist(xlTask.SavePath()))
		if err != nil {
			_ = config.Logger.Warn(err)
			return nil
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
				return nil
			}
			if existTempFile {
				return nil
			}
		}

		return []*dto.XLTask{xlTask}
	})
	if len(downTasks) == 0 {
		return nil
	}

	err = Completed(ctx, downTasks...)
	if err != nil {
		return err
	}

	return nil
}

// Completed 下载完成
func Completed(ctx context.Context, tasks ...*dto.XLTask) error {
	_ = config.Logger.Infof("tasks `%v` download complete", stlslices.Map(tasks, func(_ int, t *dto.XLTask) string { return t.Name() }))

	// 查找种子文件
	hash2TaskAndTorrent := stlslices.ToMap(tasks, func(t *dto.XLTask) (string, tuple.Tuple3[*dto.XLTask, string, *metainfo.Info]) {
		return t.Hash(), tuple.Pack3[*dto.XLTask, string, *metainfo.Info](t, "", nil)
	})
	fileEntries, err := stlerr.ErrorWith(os.ReadDir(config.XLBtDir))
	if err != nil {
		return err
	}
	for _, fileEntry := range fileEntries {
		fp := filepath.Join(config.XLBtDir, fileEntry.Name())
		tmi, err := stlerr.ErrorWith(metainfo.LoadFromFile(fp))
		if err != nil {
			_ = config.Logger.Debug(err)
			continue
		}
		hash := tmi.HashInfoBytes().HexString()
		tt, ok := hash2TaskAndTorrent[hash]
		if !ok {
			continue
		}
		tmiInfo, err := stlerr.ErrorWith(tmi.UnmarshalInfo())
		if err != nil {
			_ = config.Logger.Debug(err)
			continue
		}
		hash2TaskAndTorrent[hash] = tuple.Pack3(tt.E1(), fp, &tmiInfo)
	}

	// 迁移
	for _, taskAndTorrentPath := range hash2TaskAndTorrent {
		task, tp, tmi := taskAndTorrentPath.Unpack()
		if tp == "" {
			_ = config.Logger.Warnf("can not find torrent for task `%s`", task.Name())
			continue
		}

		// 迁移文件
		exist, err := stlerr.ErrorWith(stlos.Exist(task.SavePath()))
		if err != nil {
			return err
		}
		if exist {
			fromPath := task.SavePath()
			toPath := filepath.Join(config.DownloadCompleteDir, tmi.Name)
			singleFile := len(tmi.Files) == 1
			if singleFile {
				// 单文件时迅雷会自动创建一个文件夹
				fromPath = filepath.Join(fromPath, tmi.Name)
			}
			err = stlerr.ErrorWrap(os.Rename(fromPath, toPath))
			if err != nil {
				_ = config.Logger.Warnf("tmi=%s, fromPath=%s, toPath=%s, xltask=%s", util.ToJson[string](tmi), fromPath, toPath, util.ToJson[string](task.TaskInfo))
				return err
			}
		}

		// 新建qb任务
		var existedQbTask *dto.QBTask
		if existedTask := stlval.IgnoreWith(GetTaskByHash(ctx, task.Hash())); existedTask != nil {
			existedQbTask, _ = existedTask.(*dto.QBTask)
		}
		if existedQbTask == nil {
			err = stlerr.ErrorWrap(qb.Client.AddTorrentFromFileCtx(ctx, tp, map[string]string{
				"skip_checking": "true",
				"rename":        task.Name(),
			}))
			if err != nil {
				return err
			}
		}

		// 删除迅雷任务
		err = stlerr.ErrorWrap(xl.Client.DeleteTask(ctx, task.TaskInfo.ID, false))
		if err != nil {
			return err
		}

		// 迁移种子文件
		err = stlerr.ErrorWrap(os.Rename(tp, filepath.Join(config.TorrentDir, task.Hash()+".torrent")))
		if err != nil {
			_ = config.Logger.Warn(err)
		}

		// 回调
		if config.TaskDownloadCompletedFallbackAddr != "" {
			fallbackReq := &vo.TaskDownloadCompletedFallbackRequest{
				Name:     task.Name(),
				Hash:     task.Hash(),
				SavePath: task.SavePath(),
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
