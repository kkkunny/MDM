package dto

import (
	"fmt"
	"regexp"
	"time"

	"github.com/autobrr/go-qbittorrent"
	stlslices "github.com/kkkunny/stl/container/slices"
	stlos "github.com/kkkunny/stl/os"
	"github.com/kkkunny/xunlei/dto"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/dal/db/po"
	"github.com/kkkunny/MDM/model/vo"
)

// Task 任务信息
type Task interface {
	ID() string
	Hash() string
	Name() string
	Category() string
	SavePath() string
	Phase() vo.TaskPhase
	Speed() stlos.Size
	CreatedAt() time.Time
	ToVO() *vo.Task
	SetDB(task *po.Task)
}

const XLTaskIDPrefix = "XL|"

type XLTask struct {
	task *po.Task
	*dto.TaskInfo
}

func (t XLTask) ID() string {
	if t.task != nil {
		return *t.task.ID
	}
	return fmt.Sprintf("%s%s", XLTaskIDPrefix, t.TaskInfo.ID)
}

func (t XLTask) SavePath() string {
	path, _ := stlos.ReplaceBase(t.TaskInfo.SavePath, config.XLDownloadDirByXL, config.DownloadDir)
	return path
}

func TaskFromXL(t *dto.TaskInfo) Task {
	return &XLTask{TaskInfo: t}
}

func (t XLTask) Hash() string {
	return t.TaskInfo.Extra["info_hash"]
}

func (t XLTask) Name() string {
	categoryMatches := xlTaskCategoryMatch.FindAllStringSubmatch(t.TaskInfo.Name, -1)
	if len(categoryMatches) > 0 {
		return categoryMatches[0][2]
	}
	return t.TaskInfo.Name
}

func (t XLTask) Category() string {
	categoryMatches := xlTaskCategoryMatch.FindAllStringSubmatch(t.TaskInfo.Name, -1)
	if len(categoryMatches) > 0 {
		return categoryMatches[0][1]
	}
	return ""
}

func (t XLTask) Phase() vo.TaskPhase {
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

func (t XLTask) Speed() stlos.Size {
	return stlos.Size(t.TaskInfo.Speed) * stlos.Byte
}

func (t XLTask) CreatedAt() time.Time {
	return t.TaskInfo.CreatedTime
}

var xlTaskCategoryMatch = regexp.MustCompile(`\[\[(.*?)]]\|(.+)`)

func (t XLTask) ToVO() *vo.Task {
	vt := &vo.Task{
		Id:        t.ID(),
		Name:      t.Name(),
		Phase:     t.Phase(),
		Size:      uint64(t.FileSize),
		CreatedAt: uint64(t.CreatedAt().Unix()),
	}
	if category := t.Category(); category != "" {
		vt.Category = &vo.Category{
			Name: category,
		}
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

func (t *XLTask) SetDB(task *po.Task) {
	t.task = task
}

const QBTaskIDPrefix = "QB|"

type QBTask struct {
	task *po.Task
	*qbittorrent.Torrent
}

func (t QBTask) ID() string {
	if t.task != nil {
		return *t.task.ID
	}
	return fmt.Sprintf("%s%s", QBTaskIDPrefix, t.Hash())
}

func (t QBTask) SavePath() string {
	path, _ := stlos.ReplaceBase(t.Torrent.ContentPath, config.QBDownloadDirByQB, config.DownloadDir)
	return path
}

func TaskFromQB(t *qbittorrent.Torrent) Task {
	return &QBTask{Torrent: t}
}

func (t QBTask) Hash() string {
	return t.Torrent.Hash
}

func (t QBTask) Name() string {
	return t.Torrent.Name
}

func (t QBTask) Category() string {
	return t.Torrent.Category
}

func (t QBTask) Phase() vo.TaskPhase {
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
		return vo.TaskPhase_TpUpRunning
	case qbittorrent.TorrentStatePausedUp:
		return vo.TaskPhase_TpUpPaused
	case qbittorrent.TorrentStateStoppedUp:
		if t.NumComplete > 0 {
			return vo.TaskPhase_TpUpCompleted
		}
		return vo.TaskPhase_TpUpPaused
	case qbittorrent.TorrentStateMissingFiles:
		return vo.TaskPhase_TpUpFailed
	default:
		return vo.TaskPhase_TpUnknown
	}
}

func (t QBTask) Speed() stlos.Size {
	if t.DlSpeed != 0 {
		return stlos.Size(t.DlSpeed) * stlos.Byte
	}
	return stlos.Size(t.UpSpeed) * stlos.Byte
}

func (t QBTask) CreatedAt() time.Time {
	return time.Unix(t.AddedOn, 0)
}

func (t QBTask) ToVO() *vo.Task {
	vt := &vo.Task{
		Id:        t.ID(),
		Name:      t.Name(),
		Phase:     t.Phase(),
		Size:      uint64(t.Size),
		CreatedAt: uint64(t.CreatedAt().Unix()),
	}
	if category := t.Category(); category != "" {
		vt.Category = &vo.Category{
			Name: category,
		}
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

func (t *QBTask) SetDB(task *po.Task) {
	t.task = task
}
