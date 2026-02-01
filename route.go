package main

import (
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/handler"
)

func Route(root *echo.Group) {
	root.GET("/ping", handler.Ping)
}
