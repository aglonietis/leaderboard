package repositories

import (
	"leaderboard/internal/database"
	"leaderboard/internal/models"
	"time"
)

type scoreRepository struct{}

type ScoreRepository interface {
	FindByName(name string) (models.PlayerScore, error)
	Create(name string) models.PlayerScore
	Save(s models.PlayerScore) (models.PlayerScore, error)
}

func NewScoreRepository() ScoreRepository {
	return &scoreRepository{}
}

func (r *scoreRepository) Save(s models.PlayerScore) (models.PlayerScore, error) {
	db := database.DbManager()

	err := db.Save(&s).Error

	return s, err
}

func (r *scoreRepository) Create(name string) models.PlayerScore {
	return models.PlayerScore{
		Name:       name,
		Score:      0,
		SubmittedAt: time.Now(),
	}
}

func (r *scoreRepository) FindByName(name string) (models.PlayerScore, error) {
	var existingScore models.PlayerScore

	db := database.DbManager()

	err := db.Where("name = ?", name).First(&existingScore).Error

	return existingScore, err
}
