package handler

import (
	"net/http"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	xldto "github.com/kkkunny/xunlei/dto"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/dal/xunlei"
	"github.com/kkkunny/MDM/model/dto"
	"github.com/kkkunny/MDM/model/vo"
)

func TaskList(c *echo.Context) error {
	ctx := c.Request().Context()
	xlTasks, err := stlerr.ErrorWith(xunlei.Client.ListTasks(ctx))
	if err != nil {
		return err
	}
	tasks := stlslices.Map(xlTasks, func(_ int, xlTask *xldto.TaskInfo) *vo.Task {
		return dto.TaskFromXunlei(xlTask).VO()
	})
	return c.JSON(http.StatusOK, tasks)
}
