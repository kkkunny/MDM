package handler

import (
	"context"
	"net/http"
	"strings"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	stlval "github.com/kkkunny/stl/value"
	xldto "github.com/kkkunny/xunlei/dto"
	"github.com/labstack/echo/v5"
	"golang.org/x/sync/errgroup"

	"github.com/kkkunny/MDM/dal/db"
	"github.com/kkkunny/MDM/dal/qb"
	"github.com/kkkunny/MDM/dal/xl"
	"github.com/kkkunny/MDM/model/dto"
	"github.com/kkkunny/MDM/model/vo"
	taskSvr "github.com/kkkunny/MDM/service/task"
	"github.com/kkkunny/MDM/util"
)

func OperateTasks(c *echo.Context) error {
	ctx := c.Request().Context()

	var req vo.OperateTasksRequest
	err := stlerr.ErrorWrap(c.Bind(&req))
	if err != nil {
		return util.NewHttpError(http.StatusBadRequest, err)
	}

	if !stlslices.Contain([]vo.Operate{
		vo.Operate_OpDelete,
		vo.Operate_OpResume,
		vo.Operate_OpPause,
		vo.Operate_OpRetry,
	}, req.GetOperate()) || len(req.Ids) == 0 {
		return util.NewHttpError(http.StatusBadRequest, stlerr.Errorf("params invalid"))
	}

	ids := req.GetIds()
	d, err := db.NewTasksDal(ctx)
	if err != nil {
		return err
	}
	tasks, err := d.MGetByIDs(ids...)
	if err != nil {
		return err
	}
	for i, id := range ids {
		t, ok := tasks[id]
		if !ok {
			continue
		}
		if stlval.DerefPtrOr(t.Xlid, "") != "" {
			ids[i] = dto.XLTaskIDPrefix + *t.Xlid
		}
		if stlval.DerefPtrOr(t.Qbid, "") != "" {
			ids[i] = dto.QBTaskIDPrefix + *t.Qbid
		}
	}
	err = operateTasksByRelIDs(ctx, req.GetOperate(), ids...)
	if err != nil {
		return err
	}

	taskSvr.NotifyTasksChanged()

	return c.String(http.StatusOK, http.StatusText(http.StatusOK))
}

func operateTasksByRelIDs(ctx context.Context, op vo.Operate, ids ...string) error {
	xlIDs := stlslices.Map(stlslices.Filter(ids, func(_ int, id string) bool {
		return strings.HasPrefix(id, dto.XLTaskIDPrefix)
	}), func(_ int, id string) string {
		return strings.TrimPrefix(id, dto.XLTaskIDPrefix)
	})
	qbIDs := stlslices.Map(stlslices.Filter(ids, func(_ int, id string) bool {
		return strings.HasPrefix(id, dto.QBTaskIDPrefix)
	}), func(_ int, id string) string {
		return strings.TrimPrefix(id, dto.QBTaskIDPrefix)
	})
	if len(xlIDs) == 0 && len(qbIDs) == 0 {
		return nil
	}

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return operateXLTask(egCtx, op, xlIDs...)
	})

	eg.Go(func() error {
		return operateQBTask(egCtx, op, qbIDs...)
	})

	return eg.Wait()
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
				err := stlerr.ErrorWrap(xl.Client.DeleteTask(ctx, id, true))
				if err != nil {
					return err
				}
				d, err := db.NewTasksDal(ctx)
				if err != nil {
					return err
				}
				return d.DelByIDs(dto.XLTaskIDPrefix + id)
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
				tasks, err := stlerr.ErrorWith(xl.Client.ListTasks(ctx, xldto.TaskPhaseTypeError))
				if err != nil {
					return err
				}
				task, ok := stlslices.FindFirst(tasks, func(_ int, task *xldto.TaskInfo) bool {
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
