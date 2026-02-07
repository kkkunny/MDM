package handler

import (
	"net/http"

	stlerr "github.com/kkkunny/stl/error"
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
	case vo.Operate_OpContinue:
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
	default:
		return util.NewHttpError(http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, http.StatusText(http.StatusOK))
}
