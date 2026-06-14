package middleware

import (
	"errors"

	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/util"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		config.Logger.Infof("http request ==> [%s] %s", c.Request().Method, c.Request().URL.Path)

		err := next(c)
		var he *util.HttpError
		if errors.As(err, &he) {
			config.Logger.Error(he.Unwrap())
		} else if err != nil {
			config.Logger.Error(err)
		}
		return err
	}
}
