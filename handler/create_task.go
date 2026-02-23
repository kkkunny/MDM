package handler

import (
	"fmt"
	"net/http"

	stlerr "github.com/kkkunny/stl/error"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/dal/xl"
	"github.com/kkkunny/MDM/model/vo"
	"github.com/kkkunny/MDM/util"
)

func CreateTask(c *echo.Context) error {
	ctx := c.Request().Context()

	var req vo.CreateTaskRequest
	err := stlerr.ErrorWrap(c.Bind(&req))
	if err != nil {
		return util.NewHttpError(http.StatusBadRequest, err)
	}

	name := req.GetName()
	if req.GetCategory() != "" {
		name = fmt.Sprintf("[[%s]]|%s", req.GetCategory(), req.GetName())
	}
	xlTask, err := stlerr.ErrorWith(xl.Client.CreateTask(ctx, name, req.GetLink()))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &vo.CreateTaskResponse{
		Id: xlTask.ID,
	})
}
