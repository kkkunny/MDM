package runner

import (
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"os"
	"time"

	stlerr "github.com/kkkunny/stl/error"
	stlval "github.com/kkkunny/stl/value"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/handler"
	"github.com/kkkunny/MDM/middleware"
	"github.com/kkkunny/MDM/util"
)

func init() {
	Runners = append(Runners, RunHttp)
}

func RunHttp() (<-chan struct{}, <-chan error) {
	svr := echo.New()
	svr.Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(math.MaxInt)}))
	svr.Validator = util.NewValidator()

	route(svr.Group(""))

	succCh := make(chan struct{})
	go func() {
		defer close(succCh)
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for _ = range t.C {
			ok := func() bool {
				resp, err := http.Get(fmt.Sprintf("http://localhost%s/api/ping", stlval.Ternary(config.Release, "", ":8080")))
				if err != nil {
					return false
				}
				defer resp.Body.Close()
				if resp.StatusCode != http.StatusOK {
					return false
				}
				return true
			}()
			if !ok {
				continue
			}
			return
		}
	}()

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		errCh <- stlerr.ErrorWrap(svr.Start(fmt.Sprintf(":%d", stlval.Ternary(config.Release, 80, 8080))))
	}()

	return succCh, errCh
}

func route(root *echo.Group) {
	root.Use(
		middleware.Response,
		middleware.Logger,
		middleware.Recover,
	)

	root.Static("/", "static")

	api := root.Group("/api")
	{
		api.GET("/ping", handler.Ping)

		task := api.Group("/task")
		{
			task.GET("/list", handler.ListTasks)
			task.POST("/create", handler.CreateTask)
			task.POST("/operate", handler.OperateTasks)
		}
	}
}
