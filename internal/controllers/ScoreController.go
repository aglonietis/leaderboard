package controllers

import (
	"github.com/labstack/echo/v4"
	"leaderboard/internal/handler"
	"leaderboard/internal/repositories"
	"leaderboard/internal/requests"
	"net/http"
	"time"
)

type scoreController struct {
	scoreRepository repositories.ScoreRepository
}

type ScoreController interface {
	Store(ctx echo.Context) error
}

func NewScoreController() ScoreController {
	return &scoreController{
		scoreRepository: repositories.NewScoreRepository(),
	}
}

func (c *scoreController) Store(ctx echo.Context) error {
	var request requests.ScoreStoreRequest

	err := handler.BindValidate(ctx, &request)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	s, err := c.scoreRepository.FindByName(request.Name)

	if err != nil {
		s = c.scoreRepository.Create(request.Name)
	}

	if request.Score > s.Score {
		s.Score = request.Score
		s.SubmitedAt = time.Now()
	}

	s, err = c.scoreRepository.Save(s)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update Player Score")
	}

	return ctx.JSON(http.StatusOK,s)
}