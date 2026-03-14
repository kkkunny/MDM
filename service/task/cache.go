package task

import (
	"context"
	"time"

	"github.com/autobrr/go-qbittorrent"
	stlmaps "github.com/kkkunny/stl/container/maps"
	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	stlsync "github.com/kkkunny/stl/sync"
	xldto "github.com/kkkunny/xunlei/dto"

	"golang.org/x/sync/errgroup"

	"github.com/kkkunny/MDM/dal/db"
	"github.com/kkkunny/MDM/dal/db/po"
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

	// 查询下游服务
	eg, egCtx := errgroup.WithContext(ctx)

	var xlTasks []dto.Task
	eg.Go(func() error {
		tasks, err := stlerr.ErrorWith(xl.Client.ListTasks(egCtx))
		if err != nil {
			return err
		}
		xlTasks = stlslices.Map(tasks, func(_ int, xlt *xldto.TaskInfo) dto.Task {
			return dto.TaskFromXL(xlt)
		})
		return nil
	})

	var qbTasks []dto.Task
	eg.Go(func() error {
		tasks, err := stlerr.ErrorWith(qb.Client.GetTorrentsCtx(egCtx, qbittorrent.TorrentFilterOptions{}))
		if err != nil {
			return err
		}
		qbTasks = stlslices.Map(tasks, func(_ int, qbt qbittorrent.Torrent) dto.Task {
			return dto.TaskFromQB(&qbt)
		})
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	tasks := stlslices.Union(xlTasks, qbTasks)

	// 查询数据库
	d, err := db.NewTasksDal(ctx)
	if err != nil {
		return nil, err
	}
	dbTasks, err := d.MGetByIDs(stlslices.Map(tasks, func(_ int, t dto.Task) string { return t.ID() })...)
	if err != nil {
		return nil, err
	}

	// 回填
	for _, t := range tasks {
		id := t.ID()
		qbTask, ok := dbTasks[id]
		if !ok {
			continue
		}
		t.SetDB(qbTask)
		delete(dbTasks, id)
	}

	// 删除冗余数据库数据
	needDelIDs := stlmaps.ToSlice(dbTasks, func(_ string, t *po.Task) string { return *t.ID })
	if err = stlerr.ErrorWrap(d.DelByIDs(needDelIDs...)); err != nil {
		return nil, err
	}

	tc.data = tasks
	tc.updateAt = time.Now()
	return tc.data, nil
}
