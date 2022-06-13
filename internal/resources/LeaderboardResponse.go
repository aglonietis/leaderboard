package resources

import (
	"leaderboard/internal/database"
	"leaderboard/internal/models"
)

type LeaderboardResource struct {
	Results  []models.PlayerScore `json:"results"`
	AroundMe []models.PlayerScore `json:"around_me,omitempty"`
	database.Pagination
}
