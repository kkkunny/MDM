package task

import (
	"context"
	"os"
	"path/filepath"

	"github.com/anacrolix/torrent/metainfo"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/container/tuple"
	stlerr "github.com/kkkunny/stl/error"
	stlos "github.com/kkkunny/stl/os"
	stlval "github.com/kkkunny/stl/value"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/dal/qb"
	"github.com/kkkunny/MDM/dal/xl"
	"github.com/kkkunny/MDM/model/dto"
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

// DownloadCompleted 下载完成
func DownloadCompleted(ctx context.Context, tasks ...*dto.XLTask) error {
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
		existedQbTask := stlval.IgnoreWith(GetTaskByHash(ctx, task.Hash()))
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
	}
	return nil
}
