package task

import (
	"context"
	"time"

	"github.com/autobrr/go-qbittorrent"
	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	stlsync "github.com/kkkunny/stl/sync"
	xldto "github.com/kkkunny/xunlei/dto"

	"github.com/kkkunny/MDM/dal/qb"
	"github.com/kkkunny/MDM/dal/xl"
	"github.com/kkkunny/MDM/model/dto"
)

const cacheDuration = time.Second * 2

var tasksCache = &_TasksCache{lock: stlsync.NewReentrantRWLock()}

type _TasksCache struct {
	lock     *stlsync.ReentrantRWLock
	data     []dto.Task
	updateAt time.Time
}

func (tc *_TasksCache) Get(ctx context.Context) ([]dto.Task, error) {
	tasks, ok := tc.tryGet()
	if ok {
		return tasks, nil
	}

	tc.lock.Lock()
	defer tc.lock.Unlock()
	tasks, ok = tc.tryGet()
	if ok {
		return tasks, nil
	}

	return tc.GetLatest(ctx)
}

func (tc *_TasksCache) tryGet() ([]dto.Task, bool) {
	tc.lock.RLock()
	defer tc.lock.RUnlock()

	if time.Since(tc.updateAt) > cacheDuration {
		return nil, false
	}
	return tc.data, true
}

func (tc *_TasksCache) GetLatest(ctx context.Context) ([]dto.Task, error) {
	tc.lock.Lock()
	defer tc.lock.Unlock()

	xlTasks, err := stlerr.ErrorWith(xl.Client.ListTasks(ctx))
	if err != nil {
		return nil, err
	}
	tasks := stlslices.Map(xlTasks, func(_ int, xlt *xldto.TaskInfo) dto.Task {
		return dto.TaskFromXL(xlt)
	})

	qbTasks, err := stlerr.ErrorWith(qb.Client.GetTorrentsCtx(ctx, qbittorrent.TorrentFilterOptions{}))
	if err != nil {
		return nil, err
	}
	tasks = append(tasks, stlslices.Map(qbTasks, func(_ int, qbt qbittorrent.Torrent) dto.Task {
		return dto.TaskFromQB(&qbt)
	})...)

	tc.data = tasks
	tc.updateAt = time.Now()
	return tc.data, nil
}
