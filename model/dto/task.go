package dto

import (
	"fmt"
	"regexp"
	"time"

	"github.com/autobrr/go-qbittorrent"
	stlslices "github.com/kkkunny/stl/container/slices"
	"github.com/kkkunny/xunlei/dto"

	"github.com/kkkunny/MDM/model/vo"
	"github.com/kkkunny/MDM/util"
)

// Task 任务信息
type Task interface {
	CreatedAt() time.Time
	ToVO() *vo.Task
}

type xlTask struct {
	*dto.TaskInfo
}

func TaskFromXL(t *dto.TaskInfo) Task {
	return &xlTask{TaskInfo: t}
}

func (t xlTask) CreatedAt() time.Time {
	return t.TaskInfo.CreatedTime
}

func covertXLPhase2VO(phase dto.TaskPhase) vo.TaskPhase {
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

var xlTaskCategoryMatch = regexp.MustCompile(`\[\[(.*?)]]\|(.+)`)

func (t xlTask) ToVO() *vo.Task {
	vt := &vo.Task{
		Id:        fmt.Sprintf("XL|%s", t.ID),
		Phase:     covertXLPhase2VO(t.Phase),
		Size:      uint64(t.FileSize),
		CreatedAt: uint64(t.CreatedAt().Unix()),
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
		dto.TaskPhaseTypeComplete,
		dto.TaskPhaseTypePaused,
		dto.TaskPhaseTypeError,
	}, t.Phase) {
		vt.DownloadStats = &vo.DownloadStats{
			Speed:    uint64(t.Speed),
			Progress: vt.Size * uint64(t.Progress) / 100,
		}
	}
	return vt
}

type qbTask struct {
	*qbittorrent.Torrent
}

func TaskFromQB(t *qbittorrent.Torrent) Task {
	return &qbTask{Torrent: t}
}

func (t qbTask) CreatedAt() time.Time {
	return time.Unix(t.AddedOn, 0)
}

func covertQBPhase2VO(state qbittorrent.TorrentState) vo.TaskPhase {
	switch state {
	case qbittorrent.TorrentStateQueuedDl:
		return vo.TaskPhase_TpDownWaiting
	case qbittorrent.TorrentStateAllocating, qbittorrent.TorrentStateDownloading, qbittorrent.TorrentStateMetaDl,
		qbittorrent.TorrentStateStalledDl, qbittorrent.TorrentStateCheckingDl, qbittorrent.TorrentStateForcedDl:
		return vo.TaskPhase_TpDownRunning
	case qbittorrent.TorrentStatePausedDl, qbittorrent.TorrentStateStoppedDl:
		return vo.TaskPhase_TpDownPaused
	case qbittorrent.TorrentStateError, qbittorrent.TorrentStateMissingFiles:
		return vo.TaskPhase_TpDownFailed
	default:
		return vo.TaskPhase_TpUnknown
	}
}

func (t qbTask) ToVO() *vo.Task {
	vt := &vo.Task{
		Id:        util.MD5(fmt.Sprintf("QB|%s", t.Hash)),
		Name:      t.Name,
		Phase:     covertQBPhase2VO(t.State),
		Size:      uint64(t.Size),
		CreatedAt: uint64(t.CreatedAt().Unix()),
	}
	if stlslices.Contain([]qbittorrent.TorrentState{
		qbittorrent.TorrentStateAllocating,
		qbittorrent.TorrentStateDownloading,
		qbittorrent.TorrentStateMetaDl,
		qbittorrent.TorrentStatePausedDl,
		qbittorrent.TorrentStateStoppedDl,
		qbittorrent.TorrentStateQueuedDl,
		qbittorrent.TorrentStateStalledDl,
		qbittorrent.TorrentStateCheckingDl,
		qbittorrent.TorrentStateForcedDl,
	}, t.State) {
		vt.DownloadStats = &vo.DownloadStats{
			Speed:    uint64(t.DlSpeed),
			Progress: uint64(t.Downloaded),
		}
	}
	return vt
}
