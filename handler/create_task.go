package handler

import (
	"net/http"

	stlerr "github.com/kkkunny/stl/error"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/dal/xunlei"
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

	// name := fmt.Sprintf("[%s]%s", req.GetCategory(), req.GetName())
	xlTask, err := stlerr.ErrorWith(xunlei.Client.CreateTask(ctx, req.GetName(), req.GetLink()))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &vo.CreateTaskResponse{
		Id: xlTask.ID,
	})
}
