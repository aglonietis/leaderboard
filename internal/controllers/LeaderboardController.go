package controllers

import (
	"github.com/labstack/echo/v4"
	"leaderboard/internal/handler"
	"leaderboard/internal/models"
	"leaderboard/internal/repositories"
	"leaderboard/internal/resources"
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

	page := handler.QueryParamInt(ctx, "page", 1)

	scores, pagination, err := c.leaderboardRepository.Index(page)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to retrieve Player Score List")
	}

	var aroundMe []models.PlayerScore

	playerName := ctx.QueryParam("name")

	if playerName != "" && false == containsPlayerName(scores, playerName) {
		aroundMe, err = c.leaderboardRepository.GetAroundPlayer(playerName)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, "Failed to retrieve Player Score List around player")
		}
	}

	resource := resources.LeaderboardResource{
		Pagination: pagination,
		Results:    scores,
		AroundMe:   aroundMe,
	}

	return ctx.JSON(http.StatusOK, resource)
}

// TODO: Manual implementation until "leaderboard/internal/handler" ContainsParam is fixed.
func containsPlayerName(elements []models.PlayerScore, name string) bool {
	for _, element := range elements {
		if element.Name == name {
			return true
		}
	}

	return false
}
