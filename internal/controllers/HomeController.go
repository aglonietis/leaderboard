package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type homeController struct {}

type HomeController interface {
	Index(ctx echo.Context) error
}

func NewHomeController() HomeController {
	return &homeController{}
}

func (c *homeController) Index(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
