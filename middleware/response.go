package middleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/util"
)

func Response(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		if err := next(c); err != nil {
			var he *util.HttpError
			if errors.As(err, &he) {
				_ = c.String(he.Code(), http.StatusText(he.Code()))
				return nil
			}
			_ = c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		return nil
	}
}
