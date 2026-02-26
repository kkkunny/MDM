package handler

import (
	"net/http"

	stlos "github.com/kkkunny/stl/os"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/model/vo"
	"github.com/kkkunny/MDM/service/task"
)

func StatInfo(c *echo.Context) error {
	ctx := c.Request().Context()

	tasks, err := task.GetTasks(ctx)
	if err != nil {
		return err
	}

	resp := &vo.StatInfoResponse{
		TaskCount: uint64(len(tasks)),
	}

	for _, t := range tasks {
		switch t.Phase() {
		case vo.TaskPhase_TpDownQueued,
			vo.TaskPhase_TpDownRunning,
			vo.TaskPhase_TpDownPaused,
			vo.TaskPhase_TpDownFailed:
			resp.DlCount++
			resp.DlSpeed += uint64(t.Speed() / stlos.Byte)
		case vo.TaskPhase_TpUpQueued,
			vo.TaskPhase_TpUpRunning,
			vo.TaskPhase_TpUpPaused,
			vo.TaskPhase_TpUpFailed:
			resp.UlCount++
			resp.UlSpeed += uint64(t.Speed() / stlos.Byte)
		}
	}

	return c.JSON(http.StatusOK, resp)
}
