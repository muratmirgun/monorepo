package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Validator instance
var validate = validator.New()

func BindAndValidate(e echo.Context, i interface{}) error {
	if err := e.Bind(i); err != nil {
		return err
	}

	if err := validate.Struct(i); err != nil {
		return err
	}

	return nil
}
