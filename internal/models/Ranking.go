package models

type Ranking struct {
	Name        string    `json:"name" gorm:"migrate"`
	Score       int       `json:"score" gorm:"migrate"`
	Rank        int       `json:"rank" gorm:"migrate"`
}
