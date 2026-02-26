package handler

import (
	"cmp"
	"net/http"
	"slices"

	"github.com/autobrr/go-qbittorrent"
	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	xldto "github.com/kkkunny/xunlei/dto"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/dal/qb"
	"github.com/kkkunny/MDM/dal/xl"
	"github.com/kkkunny/MDM/model/dto"
	"github.com/kkkunny/MDM/model/vo"
)

func ListTasks(c *echo.Context) error {
	ctx := c.Request().Context()

	xlTasks, err := stlerr.ErrorWith(xl.Client.ListTasks(ctx))
	if err != nil {
		return err
	}
	tasks := stlslices.Map(xlTasks, func(_ int, xlt *xldto.TaskInfo) dto.Task {
		return dto.TaskFromXL(xlt)
	})

	qbTasks, err := stlerr.ErrorWith(qb.Client.GetTorrentsCtx(ctx, qbittorrent.TorrentFilterOptions{}))
	if err != nil {
		return err
	}
	tasks = append(tasks, stlslices.Map(qbTasks, func(_ int, qbt qbittorrent.Torrent) dto.Task {
		return dto.TaskFromQB(&qbt)
	})...)

	slices.SortFunc(tasks, func(i, j dto.Task) int {
		return -cmp.Compare(i.CreatedAt().UnixNano(), j.CreatedAt().UnixNano())
	})

	return c.JSON(http.StatusOK, &vo.ListTasksResponse{
		Tasks: stlslices.Map(tasks, func(_ int, t dto.Task) *vo.Task { return t.ToVO() }),
	})
}
