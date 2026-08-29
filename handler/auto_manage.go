package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
	taskSvr "github.com/kkkunny/MDM/service/task"
)

func AutoManage(c *echo.Context) error {
	ctx := c.Request().Context()

	go func() {
		err := taskSvr.AutoManageTasks(context.WithoutCancel(ctx))
		if err != nil {
			_ = config.Logger.Error(err)
		}
		taskSvr.NotifyTasksChanged()
	}()

	return c.String(http.StatusOK, http.StatusText(http.StatusOK))
}
