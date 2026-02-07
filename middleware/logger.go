package middleware

import (
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		config.HttpLogger.Infof("http request ==> [%s] %s", c.Request().Method, c.Path())

		err := next(c)
		if err != nil {
			config.HttpLogger.Error(err)
		}
		return err
	}
}
