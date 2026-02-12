package handler

import (
	"net/http"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	"github.com/kkkunny/xunlei/dto"
	"github.com/labstack/echo/v5"
	"golang.org/x/sync/errgroup"

	"github.com/kkkunny/MDM/dal/xunlei"
	"github.com/kkkunny/MDM/model/vo"
	"github.com/kkkunny/MDM/util"
)

func OperateTasks(c *echo.Context) error {
	ctx := c.Request().Context()

	var req vo.OperateTasksRequest
	err := stlerr.ErrorWrap(c.Bind(&req))
	if err != nil {
		return util.NewHttpError(http.StatusBadRequest, err)
	}

	switch req.GetOperate() {
	case vo.Operate_OpDelete:
		var eg errgroup.Group
		for _, id := range req.GetIds() {
			eg.Go(func() error {
				return stlerr.ErrorWrap(xunlei.Client.DeleteTask(ctx, id, true))
			})
		}
		err = eg.Wait()
		if err != nil {
			return err
		}
	case vo.Operate_OpResume:
		var eg errgroup.Group
		for _, id := range req.GetIds() {
			eg.Go(func() error {
				return stlerr.ErrorWrap(xunlei.Client.ContinueTask(ctx, id))
			})
		}
		err = eg.Wait()
		if err != nil {
			return err
		}
	case vo.Operate_OpPause:
		var eg errgroup.Group
		for _, id := range req.GetIds() {
			eg.Go(func() error {
				return stlerr.ErrorWrap(xunlei.Client.PauseTask(ctx, id))
			})
		}
		err = eg.Wait()
		if err != nil {
			return err
		}
	case vo.Operate_OpRetry:
		var eg errgroup.Group
		for _, id := range req.GetIds() {
			eg.Go(func() error {
				tasks, err := stlerr.ErrorWith(xunlei.Client.ListTasks(ctx, dto.TaskPhaseTypeError))
				if err != nil {
					return err
				}
				task, ok := stlslices.FindFirst(tasks, func(_ int, task *dto.TaskInfo) bool {
					return task.ID == id
				})
				if !ok {
					return stlerr.Errorf("task not found, id=%s", id)
				}
				err = xunlei.Client.DeleteTask(ctx, task.ID, false)
				if err != nil {
					return err
				}
				_, err = stlerr.ErrorWith(xunlei.Client.CreateTask(ctx, task.Name, task.URL))
				return err
			})
		}
		err = eg.Wait()
		if err != nil {
			return err
		}
	default:
		return stlerr.Errorf("unknown operate")
	}
	return c.String(http.StatusOK, http.StatusText(http.StatusOK))
}
