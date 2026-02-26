package handler

import (
	"cmp"
	"net/http"
	"slices"

	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/model/dto"
	"github.com/kkkunny/MDM/model/vo"
	"github.com/kkkunny/MDM/service/task"
)

func ListTasks(c *echo.Context) error {
	ctx := c.Request().Context()

	tasks, err := task.GetTasks(ctx)
	if err != nil {
		return err
	}

	slices.SortFunc(tasks, func(i, j dto.Task) int {
		return -cmp.Compare(i.CreatedAt().UnixNano(), j.CreatedAt().UnixNano())
	})

	return c.JSON(http.StatusOK, &vo.ListTasksResponse{
		Tasks: stlslices.Map(tasks, func(_ int, t dto.Task) *vo.Task { return t.ToVO() }),
	})
}
