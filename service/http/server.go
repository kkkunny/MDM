package http

import (
	stlerr "github.com/kkkunny/stl/error"
	"github.com/labstack/echo/v5"
)

func Run() error {
	svr := echo.New()

	Route(svr.Group(""))

	if err := stlerr.ErrorWrap(svr.Start(":80")); err != nil {
		return err
	}
	return nil
}
