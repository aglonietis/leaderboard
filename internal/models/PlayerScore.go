package models

import (
	"gorm.io/gorm"
	"time"
)

// PlayerScore is combination of User and Score
type PlayerScore struct {
	gorm.Model  `json:"-"`
	SubmittedAt time.Time `json:"-" gorm:"index"`
	Name        string    `json:"name" gorm:"index"`
	Score       int       `json:"score" gorm:"index"`
}
