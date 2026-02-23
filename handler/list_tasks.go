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
	"github.com/kkkunny/MDM/util"
)

func ListTasks(c *echo.Context) error {
	ctx := c.Request().Context()

	var req vo.ListTasksRequest
	err := stlerr.ErrorWrap(c.Bind(&req))
	if err != nil {
		return util.NewHttpError(http.StatusBadRequest, err)
	}

	var hasMore bool

	// 从迅雷获取
	xlTasks, err := stlerr.ErrorWith(xl.Client.ListTasks(ctx))
	if err != nil {
		return err
	}
	tasks := stlslices.Map(xlTasks, func(_ int, xlt *xldto.TaskInfo) dto.Task {
		return dto.TaskFromXL(xlt)
	})
	slices.SortFunc(tasks, func(i, j dto.Task) int {
		return -cmp.Compare(i.CreatedAt().UnixNano(), j.CreatedAt().UnixNano())
	})
	hasMore = len(tasks) > int(req.GetPage())*int(req.GetCount())
	xlNeeds := max(len(tasks)-int(req.GetPage()-1)*int(req.GetCount()), 0)
	tasks = tasks[len(tasks)-xlNeeds:]

	// 从qb获取
	qbNeeds := int(req.GetCount()) - xlNeeds
	qbSkips := max(int(req.GetPage()-1)*int(req.GetCount())-len(xlTasks), 0)
	if qbNeeds > 0 || !hasMore {
		qbTasks, err := stlerr.ErrorWith(qb.Client.GetTorrentsCtx(ctx, qbittorrent.TorrentFilterOptions{
			Limit:   qbNeeds + 1,
			Offset:  qbSkips,
			Sort:    "added_on",
			Reverse: true,
		}))
		if err != nil {
			return err
		}
		tasks = append(tasks, stlslices.Map(qbTasks, func(_ int, qbt qbittorrent.Torrent) dto.Task {
			return dto.TaskFromQB(&qbt)
		})...)
		if len(tasks) > int(req.GetCount()) {
			hasMore = true
			tasks = tasks[:int(req.GetCount())]
		}
	}

	return c.JSON(http.StatusOK, &vo.ListTasksResponse{
		Tasks:   stlslices.Map(tasks, func(_ int, t dto.Task) *vo.Task { return t.ToVO() }),
		HasMore: hasMore,
	})
}
