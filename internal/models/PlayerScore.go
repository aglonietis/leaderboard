package models

import (
	"gorm.io/gorm"
	"time"
)

// PlayerScore is combination of User and Score
type PlayerScore struct {
	gorm.Model `json:"-"`
	SubmitedAt time.Time `json:"-"`
	Name       string    `json:"name"`
	Score      int       `json:"score"`
	Rank       int       `json:"rank" gorm:"-"`
}
