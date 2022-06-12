package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CutstomValidator :
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate : Validate Data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func NewValidator() *CustomValidator {
	return &CustomValidator{Validator: validator.New()}
}

func BindValidate(ctx echo.Context,i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		return err
	}
	if err := ctx.Validate(i); err != nil {
		return err
	}

	return nil
}