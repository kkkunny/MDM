package main

import (
	"fmt"
	"log/slog"
	"math"
	"os"

	stlerr "github.com/kkkunny/stl/error"
	stlval "github.com/kkkunny/stl/value"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/handler"
	"github.com/kkkunny/MDM/middleware"
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

		task := api.Group("/task")
		{
			task.GET("/list", handler.ListTasks)
			task.POST("/create", handler.CreateTask)
			task.POST("/operate", handler.OperateTasks)
		}
	}
}

func main() {
	svr := echo.New()
	svr.Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(math.MaxInt)}))
	svr.Validator = util.NewValidator()
	svr.JSONSerializer = new(util.ProtobufJsonEchoSerializer)

	route(svr.Group(""))

	if err := stlerr.ErrorWrap(svr.Start(fmt.Sprintf(":%d", stlval.Ternary(config.Release, 80, 8080)))); err != nil {
		config.Logger.Panic(err)
		panic(err)
	}
}
