package controllers

import (
	"github.com/labstack/echo/v4"
	"leaderboard/internal/handler"
	"leaderboard/internal/models"
	"leaderboard/internal/repositories"
	"leaderboard/internal/resources"
	"log"
	"net/http"
)

type leaderboardController struct {
	leaderboardRepository repositories.LeaderboardRepository
	scoreRepository repositories.ScoreRepository
}

type LeaderboardController interface {
	Index(ctx echo.Context) error
}

func NewLeaderboardController() LeaderboardController {
	return &leaderboardController{
		leaderboardRepository: repositories.NewLeaderboardRepository(),
		scoreRepository: repositories.NewScoreRepository(),
	}
}

func (c *leaderboardController) Index(ctx echo.Context) error {

	page,err := handler.QueryParamInt(ctx, "page", 1)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Query parameter \"Page\" contains an invalid value!")
	}

	// Can pass monthly to get mothly list
	rankingType := ctx.Param("type")

	scores, pagination, err := c.leaderboardRepository.Index(page,rankingType)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to retrieve Player Score List")
	}

	var aroundMe []models.Ranking

	playerName := ctx.QueryParam("name")

	if playerName != "" && false == containsPlayerName(scores, playerName) {

		aroundMe, err = c.leaderboardRepository.GetAroundPlayer(playerName,rankingType)

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
func containsPlayerName(elements []models.Ranking, name string) bool {
	for _, element := range elements {
		if element.Name == name {
			log.Println(element.Name, name)

			return true
		}
	}

	return false
}
