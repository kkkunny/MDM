package task

import (
	"context"

	stlslices "github.com/kkkunny/stl/container/slices"

	"github.com/kkkunny/MDM/model/dto"
)

func GetTasks(ctx context.Context, forceUpdate ...bool) ([]dto.Task, error) {
	if stlslices.Last(forceUpdate) {
		return tasksCache.GetLatest(ctx)
	}
	return tasksCache.Get(ctx)
}
