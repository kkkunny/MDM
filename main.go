package main

import (
	stlerr "github.com/kkkunny/stl/error"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
)

func main() {
	svr := echo.New()

	Route(svr.Group(""))

	if err := stlerr.ErrorWrap(svr.Start(":8080")); err != nil {
		config.Logger.Panic(err)
		panic(err)
	}
}
