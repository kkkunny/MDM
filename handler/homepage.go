package handler

import (
	"net/http"

	stlos "github.com/kkkunny/stl/os"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/model/vo"
	"github.com/kkkunny/MDM/service/task"
)

// Homepage 提供给Homepage的展示信息
func Homepage(c *echo.Context) error {
	ctx := c.Request().Context()

	tasks, err := task.GetTasks(ctx)
	if err != nil {
		return err
	}

	resp := &vo.HomepageResponse{}

	var dlSpeed, ulSpeed stlos.Size
	for _, t := range tasks {
		switch t.Phase() {
		case vo.TaskPhase_TpDownQueued,
			vo.TaskPhase_TpDownRunning,
			vo.TaskPhase_TpDownPaused,
			vo.TaskPhase_TpDownFailed:
			resp.DlCount++
			dlSpeed += t.Speed()
		case vo.TaskPhase_TpUpQueued,
			vo.TaskPhase_TpUpRunning,
			vo.TaskPhase_TpUpPaused,
			vo.TaskPhase_TpUpFailed:
			resp.UlCount++
			ulSpeed += t.Speed()
		}
	}
	resp.DlSpeed = dlSpeed.String() + "/s"
	resp.UlSpeed = ulSpeed.String() + "/s"

	return c.JSON(http.StatusOK, resp)
}
