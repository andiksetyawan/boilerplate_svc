package httpserver

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type goValidator struct {
	validator *validator.Validate
}

func NewValidator(validator *validator.Validate) echo.Validator {
	return &goValidator{validator: validator}
}

func (gv *goValidator) Validate(i interface{}) error {
	if err := gv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
