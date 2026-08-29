package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	stlval "github.com/kkkunny/stl/value"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/dal/db"
	"github.com/kkkunny/MDM/dal/db/po"
	"github.com/kkkunny/MDM/dal/xl"
	"github.com/kkkunny/MDM/model/vo"
	taskSvr "github.com/kkkunny/MDM/service/task"
	"github.com/kkkunny/MDM/service/xltorrent"
	"github.com/kkkunny/MDM/util"
)

func CreateTask(c *echo.Context) error {
	ctx := c.Request().Context()

	var req vo.CreateTaskRequest
	err := stlerr.ErrorWrap(c.Bind(&req))
	if err != nil {
		return util.NewHttpError(http.StatusBadRequest, err)
	}

	if len(req.GetLinks()) == 0 {
		return util.NewHttpError(http.StatusBadRequest, stlerr.Errorf("need a link"))
	}

	name := req.GetName()
	if req.GetCategory() != "" {
		name = fmt.Sprintf("[[%s]]|%s", req.GetCategory(), req.GetName())
	}
	xlTask, err := stlerr.ErrorWith(xl.Client.CreateTask(ctx, name, stlslices.First(req.GetLinks())))
	if err != nil {
		return err
	}

	id := uuid.NewString()
	d, err := db.NewTasksDal(ctx)
	if err != nil {
		return err
	}
	if err = d.MSave(&po.Task{
		ID:             &id,
		Xlid:           &xlTask.ID,
		AvailableLinks: stlval.Ptr(util.ToJson[string](req.GetLinks())),
	}); err != nil {
		_ = xl.Client.DeleteTask(ctx, xlTask.ID, true)
		return err
	}

	go func() {
		if err = xltorrent.TorrentsCache.Scan(); err != nil {
			_ = config.Logger.Warn(err)
		}
	}()

	taskSvr.NotifyTasksChanged()

	return c.JSON(http.StatusOK, &vo.CreateTaskResponse{
		Id: xlTask.ID,
	})
}
