package task

import (
	"context"
	"time"

	stlslices "github.com/kkkunny/stl/container/slices"
	"google.golang.org/protobuf/proto"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/model/dto"
	"github.com/kkkunny/MDM/model/vo"
)

// pushInterval 后端向下游拉取并推送的间隔
const pushInterval = time.Second * 2

// pushHeartbeatInterval SSE 心跳间隔
const pushHeartbeatInterval = time.Second * 15

// PushHeartbeatInterval SSE 心跳间隔
const PushHeartbeatInterval = pushHeartbeatInterval

// fullSnapshotInterval 周期性全量推送间隔，用于客户端自愈
const fullSnapshotInterval = time.Second * 30

// 通知队列，当有更新时被触发
var pushWakeCh = make(chan struct{}, 1)

// NotifyTasksChanged 唤醒推送循环，立即拉取并推送最新任务数据
func NotifyTasksChanged() {
	select {
	case pushWakeCh <- struct{}{}:
	default:
	}
}

// 上次推送的任务
var lastPushTasks []dto.Task

// StartPushLoop 启动后台推送循环，任务数据变化时广播增量给所有 SSE 订阅者
func StartPushLoop(ctx context.Context) {
	var err error
	lastPushTasks, err = GetAllTasks(ctx)
	if err != nil {
		panic(err)
	}
	lastFullAt := time.Now()

	ticker := time.NewTicker(pushInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		case <-pushWakeCh:
		}

		latestTasks, err := GetAllTasks(ctx, true)
		if err != nil {
			_ = config.Logger.Error(err)
			continue
		}

		var event *vo.TaskEvent
		forceFull := time.Since(lastFullAt) >= fullSnapshotInterval
		if forceFull {
			event = &vo.TaskEvent{
				Type:  vo.TaskEventType_TetFull,
				Tasks: stlslices.Map(latestTasks, func(_ int, t dto.Task) *vo.Task { return t.ToVO() }),
			}
		} else {
			event, err = getDiffTaskEvent(ctx, latestTasks)
		}
		if err != nil {
			_ = config.Logger.Error(err)
			continue
		} else if event == nil {
			continue
		}

		lastPushTasks = latestTasks
		if forceFull {
			lastFullAt = time.Now()
		}
		tasksNotifier.Publish(event)
	}
}

// GetLastPushTaskEvent 获取最新一次推送的全量任务事件
func GetLastPushTaskEvent() *vo.TaskEvent {
	return &vo.TaskEvent{
		Type:  vo.TaskEventType_TetFull,
		Tasks: stlslices.Map(lastPushTasks, func(_ int, t dto.Task) *vo.Task { return t.ToVO() }),
	}
}

// 对比上次快照，产出增量事件
func getDiffTaskEvent(ctx context.Context, latestTasks []dto.Task) (*vo.TaskEvent, error) {
	latestTaskID2Tasks := stlslices.ToMap(latestTasks, func(task dto.Task) (string, dto.Task) {
		return task.ID(), task
	})
	lastTaskID2Tasks := stlslices.ToMap(lastPushTasks, func(task dto.Task) (string, dto.Task) {
		return task.ID(), task
	})

	diffIDs := make([]string, 0, len(latestTasks))
	removalIDs := make([]string, 0, len(latestTasks))
	for _, lastTask := range lastPushTasks {
		id := lastTask.ID()
		if latestTask, ok := latestTaskID2Tasks[id]; ok {
			if !proto.Equal(latestTask.ToVO(), lastTask.ToVO()) {
				diffIDs = append(diffIDs, id)
			}
		} else {
			removalIDs = append(removalIDs, id)
		}
	}
	// 新增任务
	for _, latestTask := range latestTasks {
		id := latestTask.ID()
		if _, ok := lastTaskID2Tasks[id]; !ok {
			diffIDs = append(diffIDs, id)
		}
	}

	if len(diffIDs) == 0 && len(removalIDs) == 0 {
		return nil, nil
	}

	diffTasks := stlslices.Filter(latestTasks, func(_ int, task dto.Task) bool {
		return stlslices.Contain(diffIDs, task.ID())
	})
	return &vo.TaskEvent{
		Type:       vo.TaskEventType_TetUpsert,
		Tasks:      stlslices.Map(diffTasks, func(_ int, t dto.Task) *vo.Task { return t.ToVO() }),
		RemovedIds: removalIDs,
	}, nil
}
