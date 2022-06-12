package repositories

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type scoreRepository struct {}

type ScoreRepository interface {
	Index(ctx echo.Context) error
}

func NewScoreRepository() ScoreRepository {
	return &scoreRepository{}
}

func (*scoreRepository) Index(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}