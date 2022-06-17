package models

import "gorm.io/gorm"

type JwtToken struct {
	gorm.Model  `json:"-"`
	Token        string    `json:"token"`
}
