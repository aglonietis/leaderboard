package resources

import (
	"leaderboard/internal/database"
	"leaderboard/internal/models"
)

type LeaderboardResource struct {
	Results  []models.Ranking `json:"results"`
	AroundMe []models.Ranking `json:"around_me,omitempty"`
	database.Pagination
}
