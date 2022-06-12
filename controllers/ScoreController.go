package controllers

import (
	"github.com/labstack/echo/v4"
	"leaderboard/repositories"
	"net/http"
)

type scoreController struct {
	scoreRepository repositories.ScoreRepository
}

type ScoreController interface {
	Test(ctx echo.Context) error
}

func NewScoreController() ScoreController {
	return &scoreController{
		scoreRepository: repositories.NewScoreRepository(),
	}
}

func (*scoreController) Test(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, Score!")
}

func (*scoreController) find(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello")
}

func (*scoreController) store(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello")
}

func (*scoreController) update(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello")
}