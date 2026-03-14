package db

import (
	"context"
	"strings"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	stlval "github.com/kkkunny/stl/value"
	"gorm.io/gen"

	"github.com/kkkunny/MDM/dal/db/po"
	"github.com/kkkunny/MDM/dal/db/query"
	"github.com/kkkunny/MDM/model/dto"
)

type TasksDal struct {
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

func (d *TasksDal) MSave(tasks ...*po.Task) error {
	return stlerr.ErrorWrap(d.query.Save(tasks...))
}

func (d *TasksDal) findByIDs(q query.ITaskDo, ids ...string) query.ITaskDo {
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

	var conds []gen.Condition
	if len(pids) > 0 {
		conds = append(conds, query.Task.ID.In(pids...))
	}
	if len(xlIDs) > 0 {
		conds = append(conds, query.Task.Xlid.In(xlIDs...))
	}
	if len(qbIDs) > 0 {
		conds = append(conds, query.Task.Qbid.In(qbIDs...))
	}

	if len(conds) == 0 {
		return q
	} else if len(conds) == 1 {
		return q.Where(conds[0])
	} else {
		for _, c := range conds {
			q = q.Or(c)
		}
		return q
	}
}

func (d *TasksDal) MGetByIDs(ids ...string) (map[string]*po.Task, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	tasks, err := stlerr.ErrorWith(d.findByIDs(d.query, ids...).Find())
	if err != nil {
		return nil, err
	}
	id2Tasks, xlid2Tasks, qbid2Tasks := make(map[string]*po.Task), make(map[string]*po.Task), make(map[string]*po.Task)
	for _, t := range tasks {
		id2Tasks[*t.ID] = t
		if stlval.DerefPtrOr(t.Xlid, "") != "" {
			xlid2Tasks[*t.Xlid] = t
		}
		if stlval.DerefPtrOr(t.Qbid, "") != "" {
			qbid2Tasks[*t.Qbid] = t
		}
	}
	taskMap := make(map[string]*po.Task)
	for _, id := range ids {
		relID := strings.TrimPrefix(strings.TrimPrefix(id, dto.XLTaskIDPrefix), dto.QBTaskIDPrefix)
		t, ok := id2Tasks[relID]
		if ok {
			taskMap[id] = t
			continue
		}
		t, ok = xlid2Tasks[relID]
		if ok {
			taskMap[id] = t
			continue
		}
		t, ok = qbid2Tasks[relID]
		if ok {
			taskMap[id] = t
			continue
		}
	}
	return taskMap, nil
}

func (d *TasksDal) DelByIDs(ids ...string) error {
	if len(ids) == 0 {
		return nil
	}

	_, err := stlerr.ErrorWith(d.findByIDs(d.query, ids...).Delete())
	return err
}
