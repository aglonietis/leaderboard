package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type leaderboardController struct {}

type LeaderboardController interface {
	Test(ctx echo.Context) error
}

func NewLeaderboardController() LeaderboardController {
	return &leaderboardController{}
}

func (*leaderboardController) Test(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, Leaderboard!")
}
