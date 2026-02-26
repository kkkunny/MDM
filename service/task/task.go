package task

import (
	"context"

	"github.com/kkkunny/MDM/model/dto"
)

func GetTasks(ctx context.Context) ([]dto.Task, error) {
	return tasksCache.Get(ctx)
}
