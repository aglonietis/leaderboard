package repositories

import (
	"leaderboard/internal/database"
	"leaderboard/internal/models"
)

type leaderboardRepository struct {}

type LeaderboardRepository interface {
	Index(page int) ([]models.PlayerScore, database.Pagination, error)
	GetAroundPlayer(playerName string) ([]models.PlayerScore, error)
}

func NewLeaderboardRepository() LeaderboardRepository {
	return &leaderboardRepository{}
}

func (r *leaderboardRepository) Index(page int) ([]models.PlayerScore,database.Pagination, error) {
	var scores []models.PlayerScore

	db := database.DbManager()

	baseQuery := db.Order("score desc")

	pagination, err := database.Paginate(baseQuery,&scores, page)

	if err != nil {
		return scores, pagination, err
	}

	return scores, pagination, nil
}


// TODO: Fix implementation. Does not work with AroundMe positioning
// TODO: Maybe move Ranking to database level or cache. Possibly NOT needed if can use PostgreSQL Index ranks
func (r *leaderboardRepository) FilLRanks(scores []models.PlayerScore, pageOffset int) []models.PlayerScore {
	rank := pageOffset + 1

	for index, score := range scores {
		score.Rank = rank
		scores[index] = score
		rank++
	}

	return scores
}

func (r *leaderboardRepository) GetAroundPlayer(playerName string) ([]models.PlayerScore, error) {
	var scores []models.PlayerScore

	// TODO: Implement AroundMe
	//db := database.DbManager()

	return scores, nil
}
