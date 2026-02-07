package util

import (
	"net/http"

	"github.com/go-playground/validator"
	stlerr "github.com/kkkunny/stl/error"
	"github.com/labstack/echo/v5"
)

type customValidator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &customValidator{validator: validator.New()}
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return stlerr.ErrorWrap(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return nil
}
