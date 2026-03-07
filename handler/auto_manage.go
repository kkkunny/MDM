package handler

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/anacrolix/torrent/metainfo"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/stl/container/tuple"
	stlerr "github.com/kkkunny/stl/error"
	xldto "github.com/kkkunny/xunlei/dto"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/dal/qb"
	"github.com/kkkunny/MDM/dal/xl"
	"github.com/kkkunny/MDM/model/dto"
	taskSvr "github.com/kkkunny/MDM/service/task"
	"github.com/kkkunny/MDM/util"
)

func AutoManage(c *echo.Context) error {
	ctx := c.Request().Context()

	tasks, err := taskSvr.GetTasks(ctx, true)
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
			config.HttpLogger.Warn(err)
			return nil
		}
		if existTempFile {
			return nil
		}
		return []*dto.XLTask{xlTask}
	})
	if len(downTasks) == 0 {
		return nil
	}

	// 查找种子文件
	hash2TaskAndTorrent := stlslices.ToMap(downTasks, func(t *dto.XLTask) (string, tuple.Tuple3[*dto.XLTask, string, *metainfo.Info]) {
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
			config.HttpLogger.Debug(err)
			continue
		}
		hash := tmi.HashInfoBytes().HexString()
		tt, ok := hash2TaskAndTorrent[hash]
		if !ok {
			continue
		}
		tmiInfo, err := stlerr.ErrorWith(tmi.UnmarshalInfo())
		if err != nil {
			config.HttpLogger.Debug(err)
			continue
		}
		hash2TaskAndTorrent[hash] = tuple.Pack3(tt.E1(), fp, &tmiInfo)
	}

	// 迁移
	for _, taskAndTorrentPath := range hash2TaskAndTorrent {
		task, tp, tmi := taskAndTorrentPath.Unpack()
		if tp == "" {
			config.HttpLogger.Warnf("can not find torrent for task `%s`", task.Name())
			continue
		}

		// 迁移文件
		fromPath := task.SavePath()
		toPath := filepath.Join(config.CompleteDownloadDir, tmi.Name)
		singleFile := len(tmi.Files) == 1
		if singleFile {
			// 单文件时迅雷会自动创建一个文件夹
			fromPath = filepath.Join(fromPath, tmi.Name)
		}
		err = stlerr.ErrorWrap(os.Rename(fromPath, toPath))
		if err != nil {
			config.HttpLogger.Warnf("tmi=%s, fromPath=%s, toPath=%s, xltask=%s", util.ToJson[string](tmi), fromPath, toPath, util.ToJson[string](task.TaskInfo))
			return err
		}

		// 新建qb任务
		err = stlerr.ErrorWrap(qb.Client.AddTorrentFromFileCtx(ctx, tp, map[string]string{
			"skip_checking": "true",
			"rename":        task.Name(),
		}))
		if err != nil {
			return err
		}

		// 删除迅雷任务
		err = stlerr.ErrorWrap(xl.Client.DeleteTask(ctx, task.TaskInfo.ID, false))
		if err != nil {
			return err
		}
	}

	return c.String(http.StatusOK, http.StatusText(http.StatusOK))
}
