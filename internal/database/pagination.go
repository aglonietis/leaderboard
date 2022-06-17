package database

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"math"
)

type Pagination struct {
	Page     int `json:"-"`
	NextPage int `json:"next_page"`
}

func Paginate(db *gorm.DB, models interface{}, page int) (Pagination, error) {
	var totalRows int64
	var pagination Pagination
	pagination.Page = page

	// Get totalPages
	db.Model(models).Count(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	if totalPages > page {
		pagination.NextPage = page + 1
	}

	// Paginate
	err := db.
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Find(models).Error

	if err != nil {
		return pagination, err
	}

	return pagination, nil
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	return viper.GetInt("LEADERBOARD_PAGE_SIZE")
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func GetAroundPlayerRange() int {
	return int(viper.GetFloat64("LEADERBOARD_PAGE_SIZE") / float64(2))
}
