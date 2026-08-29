package task

import (
	"sync"

	"github.com/kkkunny/stl/container/set"

	"github.com/kkkunny/MDM/model/vo"
)

var tasksNotifier = &_TasksNotifier{
	clients: set.StdHashSetWith[chan *vo.TaskEvent](),
}

// _TasksNotifier 任务变更广播器
type _TasksNotifier struct {
	lock    sync.RWMutex
	clients set.Set[chan *vo.TaskEvent]
}

// Subscribe 订阅任务变更，返回消息通道和取消订阅函数
func (tn *_TasksNotifier) Subscribe() (<-chan *vo.TaskEvent, func()) {
	ch := make(chan *vo.TaskEvent)

	tn.lock.Lock()
	defer tn.lock.Unlock()
	tn.clients.Add(ch)

	return ch, func() {
		tn.lock.Lock()
		defer tn.lock.Unlock()
		tn.clients.Remove(ch)
	}
}

// Publish 广播任务数据，非阻塞，慢客户端丢弃旧消息
func (tn *_TasksNotifier) Publish(data *vo.TaskEvent) {
	tn.lock.RLock()
	defer tn.lock.RUnlock()

	for ch := range tn.clients.Iter() {
		select {
		case ch <- data:
		default:
			// 通道已满，丢弃旧消息后重投
			select {
			case <-ch:
			default:
			}
			select {
			case ch <- data:
			default:
			}
		}
	}
}

// SubscribeTasks 订阅任务变更广播
func SubscribeTasks() (<-chan *vo.TaskEvent, func()) {
	return tasksNotifier.Subscribe()
}
