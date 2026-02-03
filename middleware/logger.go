package middleware

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/config"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		config.HttpLogger.Infof("http request ==> [%s] %s", c.Request().Method, c.Path())

		if err := next(c); err != nil {
			config.HttpLogger.Error(err)
			_ = c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		return nil
	}
}
