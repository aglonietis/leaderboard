package controllers

import (
	"github.com/labstack/echo/v4"
	"leaderboard/internal/handler"
	"leaderboard/internal/repositories"
	"net/http"
)

type leaderboardController struct {
	leaderboardRepository repositories.LeaderboardRepository
}

type LeaderboardController interface {
	Index(ctx echo.Context) error
}

func NewLeaderboardController() LeaderboardController {
	return &leaderboardController{
		leaderboardRepository: repositories.NewLeaderboardRepository(),
	}
}

func (c *leaderboardController) Index(ctx echo.Context) error {

	page := handler.QueryParamInt(ctx,"page",1)

	scores,err :=  c.leaderboardRepository.Index(page)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to retrieve Player Score List")
	}

	return ctx.JSON(http.StatusOK, scores)
}