package dto

import (
	stlslices "github.com/kkkunny/stl/container/slices"
	stlos "github.com/kkkunny/stl/os"
	"github.com/kkkunny/xunlei/dto"

	"github.com/kkkunny/MDM/model/vo"
)

// Task 任务信息
type Task interface {
	VO() *vo.Task
}

type xlTask dto.TaskInfo

func TaskFromXunlei(t *dto.TaskInfo) Task {
	return (*xlTask)(t)
}

func covertXunleiPhase2VO(phase dto.TaskPhase) vo.TaskPhase {
	switch phase {
	case dto.TaskPhaseTypePending:
		return vo.TaskPhase_Pending
	case dto.TaskPhaseTypeRunning:
		return vo.TaskPhase_Running
	case dto.TaskPhaseTypePaused:
		return vo.TaskPhase_Paused
	case dto.TaskPhaseTypeError:
		return vo.TaskPhase_Error
	case dto.TaskPhaseTypeComplete:
		return vo.TaskPhase_Completed
	case dto.TaskPhaseTypeDelete:
		return vo.TaskPhase_Deleted
	default:
		return vo.TaskPhase_Unknown
	}
}

func (t xlTask) VO() *vo.Task {
	vt := &vo.Task{
		Id:    t.ID,
		Name:  t.Name,
		Phase: covertXunleiPhase2VO(t.Phase),
		Size:  uint64(stlos.Byte * stlos.Size(t.FileSize)),
	}
	if stlslices.Contain([]dto.TaskPhase{
		dto.TaskPhaseTypeRunning,
		dto.TaskPhaseTypePaused,
	}, t.Phase) {
		vt.Download = &vo.DownloadStatus{
			Speed:    uint64(stlos.Byte * stlos.Size(t.Speed)),
			Progress: vt.Size * uint64(t.Progress) / 100,
		}
	}
	return vt
}
