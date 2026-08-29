package handler

import (
	"net/http"
	"time"

	stlerr "github.com/kkkunny/stl/error"
	"github.com/labstack/echo/v5"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/model/vo"
	taskSvr "github.com/kkkunny/MDM/service/task"
)

// TaskEvents SSE 推送任务列表，数据变化时实时下发
func TaskEvents(c *echo.Context) error {
	w := c.Response()
	w.Header().Set(echo.HeaderContentType, "text/event-stream; charset=utf-8")
	w.Header().Set(echo.HeaderCacheControl, "no-cache")
	w.Header().Set(echo.HeaderConnection, "keep-alive")
	w.WriteHeader(http.StatusOK)

	ctx := c.Request().Context()
	flusher := http.NewResponseController(w)
	// 清除 Echo 默认的 WriteTimeout（30s），否则 SSE 长连接会被服务端关闭
	_ = flusher.SetWriteDeadline(time.Time{})

	// 连接建立即推送当前快照
	if err := writeEventData(w, taskSvr.GetLastPushTaskEvent()); err != nil {
		_ = config.Logger.Error(err)
		return err
	}
	if err := stlerr.ErrorWrap(flusher.Flush()); err != nil {
		_ = config.Logger.Error(err)
		return err
	}

	events, unsubscribe := taskSvr.SubscribeTasks()
	defer unsubscribe()

	heartbeat := time.NewTicker(taskSvr.PushHeartbeatInterval)
	defer heartbeat.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-heartbeat.C:
			if _, err := stlerr.ErrorWith(w.Write([]byte(": ping\n\n"))); err != nil {
				_ = config.Logger.Error(err)
				return err
			}
			_ = flusher.Flush()
		case event, ok := <-events:
			if !ok {
				return nil
			}
			if err := writeEventData(w, event); err != nil {
				_ = config.Logger.Error(err)
				return err
			}
			// 立即 flush，否则事件会堆积在缓冲区直到心跳才批量下发
			_ = flusher.Flush()
		}
	}
}

func writeEventData(w http.ResponseWriter, event *vo.TaskEvent) error {
	payload, err := stlerr.ErrorWith(protojson.MarshalOptions{UseProtoNames: true}.Marshal(event))
	if err != nil {
		_ = config.Logger.Error(err)
		return nil
	}

	data := make([]byte, 0, len(payload)+7)
	data = append(data, "data: "...)
	data = append(data, payload...)
	data = append(data, '\n', '\n')
	_, err = stlerr.ErrorWith(w.Write(data))
	return err
}
