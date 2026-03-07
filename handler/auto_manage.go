package handler

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	xldto "github.com/kkkunny/xunlei/dto"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/model/dto"
	taskSvr "github.com/kkkunny/MDM/service/task"
)

func AutoManage(c *echo.Context) error {
	ctx := c.Request().Context()

	tasks, err := taskSvr.GetAllTasks(ctx, true)
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
			_ = config.Logger.Warn(err)
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

	err = taskSvr.DownloadCompleted(ctx, downTasks...)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, http.StatusText(http.StatusOK))
}
