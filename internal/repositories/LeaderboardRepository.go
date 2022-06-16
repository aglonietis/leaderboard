package repositories

import (
	"leaderboard/internal/database"
	"leaderboard/internal/models"
)

type leaderboardRepository struct{}

type LeaderboardRepository interface {
	Index(page int) ([]models.Ranking, database.Pagination, error)
	GetAroundPlayer(name string) ([]models.Ranking, error)
}

func NewLeaderboardRepository() LeaderboardRepository {
	return &leaderboardRepository{}
}

func (r *leaderboardRepository) Index(page int) ([]models.Ranking, database.Pagination, error) {
	var scores []models.Ranking

	db := database.DbManager()

	baseQuery := db.Order("rankings.rank asc").
		Joins("JOIN rankings on rankings.player_score_id=player_scores.id").
		Select("rankings.rank","player_scores.name","player_scores.score").
		Table("player_scores")

	pagination, err := database.Paginate(baseQuery, &scores, page)

	if err != nil {
		return scores, pagination, err
	}

	return scores, pagination, nil
}

func (r *leaderboardRepository) GetAroundPlayer(name string) ([]models.Ranking, error) {
	var scores []models.Ranking

	db := database.DbManager()

	playerRankingSubqueryMin := db.Select("(player_ranking.rank - 5)").
		Joins("JOIN rankings as player_ranking on player_ranking.player_score_id=player_scores.id").
		Where("name = ?", name).
		Table("player_scores")

	playerRankingSubqueryMax := db.Select("(player_ranking.rank + 5)").
		Joins("JOIN rankings as player_ranking on player_ranking.player_score_id=player_scores.id").
		Where("name = ?", name).
		Table("player_scores")

	err := db.Order("rankings.rank asc").
		Where("rankings.rank > (?)",playerRankingSubqueryMin).
		Where("rankings.rank < (?)",playerRankingSubqueryMax).
		Joins("JOIN player_scores on rankings.player_score_id=player_scores.id").
		Table("rankings").
		Select("player_scores.name","player_scores.score","rankings.rank").
		Find(&scores).Error

	if err != nil {
		return scores, err
	}



	return scores, nil
}
