package repositories

import (
	"leaderboard/internal/database"
	"leaderboard/internal/models"
	"strconv"
)

type leaderboardRepository struct{}

type LeaderboardRepository interface {
	Index(page int, rankingType string) ([]models.Ranking, database.Pagination, error)
	GetAroundPlayer(name string, rankingType string) ([]models.Ranking, error)
}

func NewLeaderboardRepository() LeaderboardRepository {
	return &leaderboardRepository{}
}

func (r *leaderboardRepository) Index(page int, rankingType string) ([]models.Ranking, database.Pagination, error) {
	var scores []models.Ranking

	db := database.DbManager()

	rankingTable := getRankingTable(rankingType)

	baseQuery := db.Order("rank asc").
		Joins("JOIN "+rankingTable+" on player_score_id=player_scores.id").
		Select("rank", "player_scores.name", "player_scores.score").
		Table("player_scores")

	pagination, err := database.Paginate(baseQuery, &scores, page)

	if err != nil {
		return scores, pagination, err
	}

	return scores, pagination, nil
}

func (r *leaderboardRepository) GetAroundPlayer(name string, rankingType string) ([]models.Ranking, error) {
	var scores []models.Ranking

	db := database.DbManager()

	rankingTable := getRankingTable(rankingType)
	aroundPlayerRange := getAroundPlayerRange()

	playerRankingSubqueryMin := db.Select("(player_ranking.rank - "+aroundPlayerRange+")").
		Joins("JOIN "+rankingTable+" as player_ranking on player_ranking.player_score_id=player_scores.id").
		Where("name = ?", name).
		Table("player_scores")

	playerRankingSubqueryMax := db.Select("(player_ranking.rank + "+aroundPlayerRange+")").
		Joins("JOIN "+rankingTable+" as player_ranking on player_ranking.player_score_id=player_scores.id").
		Where("name = ?", name).
		Table("player_scores")

	err := db.Order("rank asc").
		Where("rank > (?)", playerRankingSubqueryMin).
		Where("rank < (?)", playerRankingSubqueryMax).
		Joins("JOIN player_scores on player_score_id=player_scores.id").
		Table(rankingTable).
		Select("player_scores.name", "player_scores.score", "rank").
		Find(&scores).Error

	if err != nil {
		return scores, err
	}

	return scores, nil
}

func getRankingTable(rankingType string) string {
	if rankingType == "monthly" {
		return "rankings_monthly"
	}

	return "rankings"
}

func getAroundPlayerRange() string {
	aroundRange := database.GetAroundPlayerRange()

	return strconv.Itoa(aroundRange)
}
