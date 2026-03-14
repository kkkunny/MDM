package db

import (
	"context"
	"strings"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	"gorm.io/gen"

	"github.com/kkkunny/MDM/dal/db/po"
	"github.com/kkkunny/MDM/dal/db/query"
	"github.com/kkkunny/MDM/model/dto"
)

type TasksDal struct {
	model
	query query.ITaskDo
	ctx   context.Context
}

func NewTasksDal(ctx context.Context) (*TasksDal, error) {
	queryer, err := stlerr.ErrorWith(QueryGetter())
	if err != nil {
		return nil, err
	}
	return &TasksDal{
		query: queryer.WithContext(ctx).Task,
		ctx:   ctx,
	}, nil
}

func (d *TasksDal) Create(task *po.Task) error {
	return stlerr.ErrorWrap(d.query.Create(task))
}

func (d *TasksDal) MGetByIDs(ids ...string) (map[string]*po.Task, error) {
	pids := stlslices.Filter(ids, func(_ int, id string) bool {
		return !strings.HasPrefix(id, dto.XLTaskIDPrefix) && !strings.HasPrefix(id, dto.QBTaskIDPrefix)
	})
	xlIDs := stlslices.Map(stlslices.Filter(ids, func(_ int, id string) bool {
		return strings.HasPrefix(id, dto.XLTaskIDPrefix)
	}), func(_ int, id string) string {
		return strings.TrimPrefix(id, dto.XLTaskIDPrefix)
	})
	qbIDs := stlslices.Map(stlslices.Filter(ids, func(_ int, id string) bool {
		return strings.HasPrefix(id, dto.QBTaskIDPrefix)
	}), func(_ int, id string) string {
		return strings.TrimPrefix(id, dto.QBTaskIDPrefix)
	})
	if len(pids) == 0 && len(xlIDs) == 0 && len(qbIDs) == 0 {
		return nil, nil
	}

	var conds []gen.Condition
	if len(pids) > 0 {
		conds = append(conds, d.query.Where(
			query.Task.ID.In(pids...),
		))
	}
	if len(xlIDs) > 0 {
		conds = append(conds, d.query.Where(
			query.Task.Xlid.In(xlIDs...),
		))
	}
	if len(qbIDs) > 0 {
		conds = append(conds, d.query.Where(
			query.Task.Qbid.In(qbIDs...),
		))
	}

	tasks, err := stlerr.ErrorWith(d.query.Or(conds...).Find())
	if err != nil {
		return nil, err
	}
	id2Tasks, xlid2Tasks, qbid2Tasks := make(map[string]*po.Task), make(map[string]*po.Task), make(map[string]*po.Task)
	for _, t := range tasks {
		id2Tasks[*t.ID] = t
		if t.Xlid != "" {
			xlid2Tasks[t.Xlid] = t
		}
		if t.Qbid != "" {
			qbid2Tasks[t.Qbid] = t
		}
	}
	taskMap := make(map[string]*po.Task)
	for _, id := range ids {
		t, ok := id2Tasks[id]
		if ok {
			taskMap[id] = t
			continue
		}
		t, ok = xlid2Tasks[id]
		if ok {
			taskMap[id] = t
			continue
		}
		t, ok = qbid2Tasks[id]
		if ok {
			taskMap[id] = t
			continue
		}
	}
	return taskMap, nil
}

func (d *TasksDal) DelByIDs(ids ...string) error {
	pids := stlslices.Filter(ids, func(_ int, id string) bool {
		return !strings.HasPrefix(id, dto.XLTaskIDPrefix) && !strings.HasPrefix(id, dto.QBTaskIDPrefix)
	})
	xlIDs := stlslices.Map(stlslices.Filter(ids, func(_ int, id string) bool {
		return strings.HasPrefix(id, dto.XLTaskIDPrefix)
	}), func(_ int, id string) string {
		return strings.TrimPrefix(id, dto.XLTaskIDPrefix)
	})
	qbIDs := stlslices.Map(stlslices.Filter(ids, func(_ int, id string) bool {
		return strings.HasPrefix(id, dto.QBTaskIDPrefix)
	}), func(_ int, id string) string {
		return strings.TrimPrefix(id, dto.QBTaskIDPrefix)
	})
	if len(pids) == 0 && len(xlIDs) == 0 && len(qbIDs) == 0 {
		return nil
	}

	var conds []gen.Condition
	if len(pids) > 0 {
		conds = append(conds, d.query.Where(
			query.Task.ID.In(pids...),
		))
	}
	if len(xlIDs) > 0 {
		conds = append(conds, d.query.Where(
			query.Task.Xlid.In(xlIDs...),
		))
	}
	if len(qbIDs) > 0 {
		conds = append(conds, d.query.Where(
			query.Task.Qbid.In(qbIDs...),
		))
	}

	_, err := stlerr.ErrorWith(d.query.Or(conds...).Delete())
	return err
}
