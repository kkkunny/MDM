package dto

import (
	"fmt"
	"regexp"
	"time"

	"github.com/autobrr/go-qbittorrent"
	stlslices "github.com/kkkunny/stl/container/slices"
	stlos "github.com/kkkunny/stl/os"
	"github.com/kkkunny/xunlei/dto"

	"github.com/kkkunny/MDM/model/vo"
)

// Task 任务信息
type Task interface {
	Phase() vo.TaskPhase
	Speed() stlos.Size
	CreatedAt() time.Time
	ToVO() *vo.Task
}

type xlTask struct {
	*dto.TaskInfo
}

func TaskFromXL(t *dto.TaskInfo) Task {
	return &xlTask{TaskInfo: t}
}

func (t xlTask) Phase() vo.TaskPhase {
	switch t.TaskInfo.Phase {
	case dto.TaskPhaseTypePending:
		return vo.TaskPhase_TpDownQueued
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

func (t xlTask) Speed() stlos.Size {
	return stlos.Size(t.TaskInfo.Speed) * stlos.Byte
}

func (t xlTask) CreatedAt() time.Time {
	return t.TaskInfo.CreatedTime
}

var xlTaskCategoryMatch = regexp.MustCompile(`\[\[(.*?)]]\|(.+)`)

func (t xlTask) ToVO() *vo.Task {
	vt := &vo.Task{
		Id:        fmt.Sprintf("XL|%s", t.ID),
		Phase:     t.Phase(),
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
	if stlslices.Contain([]vo.TaskPhase{
		vo.TaskPhase_TpDownQueued,
		vo.TaskPhase_TpDownRunning,
		vo.TaskPhase_TpDownPaused,
		vo.TaskPhase_TpDownFailed,
		vo.TaskPhase_TpDownCompleted,
	}, vt.Phase) {
		vt.DownloadStats = &vo.DownloadStats{
			Speed: uint64(t.TaskInfo.Speed),
			Size:  vt.Size * uint64(t.Progress) / 100,
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

func (t qbTask) Phase() vo.TaskPhase {
	switch t.State {
	case qbittorrent.TorrentStateQueuedDl:
		return vo.TaskPhase_TpDownQueued
	case qbittorrent.TorrentStateAllocating, qbittorrent.TorrentStateDownloading, qbittorrent.TorrentStateMetaDl,
		qbittorrent.TorrentStateStalledDl, qbittorrent.TorrentStateCheckingDl, qbittorrent.TorrentStateForcedDl:
		return vo.TaskPhase_TpDownRunning
	case qbittorrent.TorrentStatePausedDl, qbittorrent.TorrentStateStoppedDl:
		return vo.TaskPhase_TpDownPaused
	case qbittorrent.TorrentStateError:
		return vo.TaskPhase_TpDownFailed
	case qbittorrent.TorrentStateCheckingUp:
		return vo.TaskPhase_TpDownCompleted
	case qbittorrent.TorrentStateQueuedUp:
		return vo.TaskPhase_TpUpQueued
	case qbittorrent.TorrentStateUploading, qbittorrent.TorrentStateStalledUp, qbittorrent.TorrentStateForcedUp:
		return vo.TaskPhase_TpDownRunning
	case qbittorrent.TorrentStatePausedUp, qbittorrent.TorrentStateStoppedUp:
		return vo.TaskPhase_TpUpPaused
	case qbittorrent.TorrentStateMissingFiles:
		return vo.TaskPhase_TpUpFailed
	default:
		return vo.TaskPhase_TpUnknown
	}
}

func (t qbTask) Speed() stlos.Size {
	if t.DlSpeed != 0 {
		return stlos.Size(t.DlSpeed) * stlos.Byte
	}
	return stlos.Size(t.UpSpeed) * stlos.Byte
}

func (t qbTask) CreatedAt() time.Time {
	return time.Unix(t.AddedOn, 0)
}

func (t qbTask) ToVO() *vo.Task {
	vt := &vo.Task{
		Id:        fmt.Sprintf("QB|%s", t.Hash),
		Name:      t.Name,
		Phase:     t.Phase(),
		Size:      uint64(t.Size),
		CreatedAt: uint64(t.CreatedAt().Unix()),
	}
	if stlslices.Contain([]vo.TaskPhase{
		vo.TaskPhase_TpDownQueued,
		vo.TaskPhase_TpDownRunning,
		vo.TaskPhase_TpDownPaused,
		vo.TaskPhase_TpDownFailed,
		vo.TaskPhase_TpDownCompleted,
	}, vt.Phase) {
		vt.DownloadStats = &vo.DownloadStats{
			Speed: uint64(t.DlSpeed),
			Size:  uint64(t.Downloaded),
		}
	}
	if stlslices.Contain([]vo.TaskPhase{
		vo.TaskPhase_TpUpQueued,
		vo.TaskPhase_TpUpRunning,
		vo.TaskPhase_TpUpPaused,
		vo.TaskPhase_TpUpFailed,
		vo.TaskPhase_TpUpCompleted,
	}, vt.Phase) {
		vt.UploadStats = &vo.UploadStats{
			Speed: uint64(t.UpSpeed),
			Size:  uint64(t.Uploaded),
		}
	}
	return vt
}
