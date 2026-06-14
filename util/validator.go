package util

import (
	"net/http"

	"github.com/go-playground/validator"
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
		return NewHttpError(http.StatusBadRequest, err)
	}
	return nil
}
