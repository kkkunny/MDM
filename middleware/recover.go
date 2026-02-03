package middleware

import (
	stlerr "github.com/kkkunny/stl/error"
	"github.com/labstack/echo/v5"
)

func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) (err error) {
		defer stlerr.RecoverTo(&err)
		return next(c)
	}
}
