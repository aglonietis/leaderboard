package models

import "gorm.io/gorm"

type User struct {
	gorm.Model  `json:"-"`
	Username string
	Password string
}
