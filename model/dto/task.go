package dto

import (
	"regexp"
	"time"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlos "github.com/kkkunny/stl/os"
	"github.com/kkkunny/xunlei/dto"

	"github.com/kkkunny/MDM/model/vo"
)

// Task 任务信息
type Task interface {
	CreatedTime() time.Time
	VO() *vo.Task
}

type xlTask struct {
	*dto.TaskInfo
}

func TaskFromXunlei(t *dto.TaskInfo) Task {
	return &xlTask{TaskInfo: t}
}

func (t xlTask) CreatedTime() time.Time {
	return t.TaskInfo.CreatedTime
}

func covertXunleiPhase2VO(phase dto.TaskPhase) vo.TaskPhase {
	switch phase {
	case dto.TaskPhaseTypePending:
		return vo.TaskPhase_TpDownWaiting
	case dto.TaskPhaseTypeRunning:
		return vo.TaskPhase_TpDownRunning
	case dto.TaskPhaseTypePaused:
		return vo.TaskPhase_TpDownPaused
	case dto.TaskPhaseTypeError:
		return vo.TaskPhase_TpDownFailed
	case dto.TaskPhaseTypeComplete:
		return vo.TaskPhase_TpDownCompleted
	default:
		return vo.TaskPhase_TpUnknown
	}
}

var xlTaskCategoryMatch = regexp.MustCompile(`\[(.*?)](.+)`)

func (t xlTask) VO() *vo.Task {
	vt := &vo.Task{
		Id:        t.ID,
		Phase:     covertXunleiPhase2VO(t.Phase),
		Size:      uint64(stlos.Byte * stlos.Size(t.FileSize)),
		CreatedAt: uint64(t.CreatedTime().Unix()),
	}
	categoryMatches := xlTaskCategoryMatch.FindAllStringSubmatch(t.Name, -1)
	if len(categoryMatches) > 0 {
		vt.Category = &vo.Category{
			Name: categoryMatches[0][1],
		}
		vt.Name = categoryMatches[0][2]
	} else {
		vt.Name = t.Name
	}
	if stlslices.Contain([]dto.TaskPhase{
		dto.TaskPhaseTypePending,
		dto.TaskPhaseTypeRunning,
		dto.TaskPhaseTypePaused,
		dto.TaskPhaseTypeError,
	}, t.Phase) {
		vt.DownloadStats = &vo.DownloadStats{
			Speed:    uint64(stlos.Byte * stlos.Size(t.Speed)),
			Progress: vt.Size * uint64(t.Progress) / 100,
		}
	}
	return vt
}
