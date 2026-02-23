package handler

import (
	"context"
	"net/http"
	"strings"
	"sync"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	stlval "github.com/kkkunny/stl/value"
	"github.com/kkkunny/xunlei/dto"
	"github.com/labstack/echo/v5"
	"golang.org/x/sync/errgroup"

	"github.com/kkkunny/MDM/dal/qb"
	"github.com/kkkunny/MDM/dal/xl"
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

	xlIDs := stlslices.FlatMap(req.GetIds(), func(_ int, id string) []string {
		if !strings.HasPrefix(id, "XL|") {
			return nil
		}
		return []string{strings.TrimPrefix(id, "XL|")}
	})
	qbIDs := stlslices.FlatMap(req.GetIds(), func(_ int, id string) []string {
		if !strings.HasPrefix(id, "QB|") {
			return nil
		}
		return []string{strings.TrimPrefix(id, "QB|")}
	})

	var wg sync.WaitGroup

	var xlErr error
	wg.Go(func() {
		xlErr = operateXLTask(ctx, req.GetOperate(), xlIDs...)
	})

	var qbErr error
	wg.Go(func() {
		qbErr = operateQBTask(ctx, req.GetOperate(), qbIDs...)
	})

	wg.Wait()
	if xlErr != nil || qbErr != nil {
		return stlval.ValueOr(xlErr, qbErr)
	}
	return c.String(http.StatusOK, http.StatusText(http.StatusOK))
}

func operateXLTask(ctx context.Context, op vo.Operate, ids ...string) (err error) {
	if len(ids) == 0 {
		return nil
	}

	switch op {
	case vo.Operate_OpDelete:
		var eg errgroup.Group
		for _, id := range ids {
			eg.Go(func() error {
				return stlerr.ErrorWrap(xl.Client.DeleteTask(ctx, id, true))
			})
		}
		err = eg.Wait()
		if err != nil {
			return err
		}
	case vo.Operate_OpResume:
		var eg errgroup.Group
		for _, id := range ids {
			eg.Go(func() error {
				return stlerr.ErrorWrap(xl.Client.ContinueTask(ctx, id))
			})
		}
		err = eg.Wait()
		if err != nil {
			return err
		}
	case vo.Operate_OpPause:
		var eg errgroup.Group
		for _, id := range ids {
			eg.Go(func() error {
				return stlerr.ErrorWrap(xl.Client.PauseTask(ctx, id))
			})
		}
		err = eg.Wait()
		if err != nil {
			return err
		}
	case vo.Operate_OpRetry:
		var eg errgroup.Group
		for _, id := range ids {
			eg.Go(func() error {
				tasks, err := stlerr.ErrorWith(xl.Client.ListTasks(ctx, dto.TaskPhaseTypeError))
				if err != nil {
					return err
				}
				task, ok := stlslices.FindFirst(tasks, func(_ int, task *dto.TaskInfo) bool {
					return task.ID == id
				})
				if !ok {
					return stlerr.Errorf("task not found, id=%s", id)
				}
				err = xl.Client.DeleteTask(ctx, task.ID, false)
				if err != nil {
					return err
				}
				_, err = stlerr.ErrorWith(xl.Client.CreateTask(ctx, task.Name, task.URL))
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
	return nil
}

func operateQBTask(ctx context.Context, op vo.Operate, ids ...string) (err error) {
	if len(ids) == 0 {
		return nil
	}

	switch op {
	case vo.Operate_OpDelete:
		return stlerr.ErrorWrap(qb.Client.DeleteTorrentsCtx(ctx, ids, true))
	case vo.Operate_OpResume:
		return stlerr.ErrorWrap(qb.Client.ResumeCtx(ctx, ids))
	case vo.Operate_OpPause:
		return stlerr.ErrorWrap(qb.Client.PauseCtx(ctx, ids))
	default:
		return stlerr.Errorf("unknown operate")
	}
}
