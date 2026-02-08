package handler

import (
	"cmp"
	"net/http"
	"slices"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	xldto "github.com/kkkunny/xunlei/dto"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/dal/xunlei"
	"github.com/kkkunny/MDM/model/dto"
	"github.com/kkkunny/MDM/model/vo"
	"github.com/kkkunny/MDM/util"
)

func ListTasks(c *echo.Context) error {
	ctx := c.Request().Context()

	var req vo.ListTasksRequest
	err := stlerr.ErrorWrap(c.Bind(&req))
	if err != nil {
		return util.NewHttpError(http.StatusBadRequest, err)
	}

	xlTasks, err := stlerr.ErrorWith(xunlei.Client.ListTasks(ctx))
	if err != nil {
		return err
	}
	tasks := stlslices.Map(xlTasks, func(_ int, xlt *xldto.TaskInfo) dto.Task {
		return dto.TaskFromXunlei(xlt)
	})

	slices.SortFunc(tasks, func(i, j dto.Task) int {
		return -cmp.Compare(i.CreatedTime().UnixNano(), j.CreatedTime().UnixNano())
	})

	if index := (req.GetPage() - 1) * req.GetCount(); index < uint32(len(xlTasks)) {
		xlTasks = xlTasks[index:]
		if uint32(len(xlTasks)) > req.GetCount() {
			xlTasks = xlTasks[:req.GetCount()]
		}
	}
	return c.JSON(http.StatusOK, &vo.ListTasksResponse{
		Tasks:   stlslices.Map(tasks, func(_ int, t dto.Task) *vo.Task { return t.VO() }),
		HasMore: uint32(len(xlTasks)) > req.GetPage()*req.GetCount(),
	})
}
