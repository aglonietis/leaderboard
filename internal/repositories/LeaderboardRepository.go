package repositories

import (
	"github.com/spf13/viper"
	"leaderboard/internal/database"
	"leaderboard/internal/models"
)

type leaderboardRepository struct {}

type LeaderboardRepository interface {
	Index(page int) ([]models.PlayerScore, error)
}

func NewLeaderboardRepository() LeaderboardRepository {
	return &leaderboardRepository{}
}

func (r *leaderboardRepository) Index(page int) ([]models.PlayerScore, error) {
	db := database.DbManager()
	var scores []models.PlayerScore

	pageSize := viper.GetInt("LEADERBOARD_PAGE_SIZE")
	pageOffset := (page - 1) * pageSize

	err := db.Order("score desc").
		Offset(pageOffset).
		Limit(pageSize).
		Find(&scores).Error

	if err != nil {
		return scores, err
	}

	return r.FilLRanks(scores, pageOffset), nil
}

// TODO: Maybe move Ranking to database levelor cache
func (r *leaderboardRepository) FilLRanks(scores []models.PlayerScore, pageOffset int) []models.PlayerScore {
	var rank int = pageOffset + 1

	for index, score := range scores {
		score.Rank = rank
		scores[index] = score
		rank++
	}

	return scores
}
