package main

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"os"
	"time"

	stlerr "github.com/kkkunny/stl/error"
	stlval "github.com/kkkunny/stl/value"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/handler"
	"github.com/kkkunny/MDM/middleware"
	taskSvr "github.com/kkkunny/MDM/service/task"
	"github.com/kkkunny/MDM/util"
)

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

		api.GET("/homepage", handler.Homepage)

		task := api.Group("/task")
		{
			task.GET("/list", handler.ListTasks)
			task.POST("/create", handler.CreateTask)
			task.POST("/operate", handler.OperateTasks)
			task.POST("/auto_manage", handler.AutoManage)
		}
	}
}

func cronjob() {
	taskAutoManageTicker := time.NewTicker(time.Second * 30)
	defer taskAutoManageTicker.Stop()

	for {
		select {
		case <-taskAutoManageTicker.C:
			err := taskSvr.AutoManageTasks(context.Background())
			if err != nil {
				_ = config.Logger.Error(err)
			}
		}
	}
}

func main() {
	svr := echo.New()
	svr.Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(math.MaxInt)}))
	svr.Validator = util.NewValidator()
	svr.JSONSerializer = new(util.ProtobufJsonEchoSerializer)

	route(svr.Group(""))
	go cronjob()

	if err := stlerr.ErrorWrap(svr.Start(fmt.Sprintf(":%d", stlval.If(config.Release, 80, 8080)))); err != nil {
		config.Logger.Panic(err)
		panic(err)
	}
}
