package repositories

import (
	"golang.org/x/crypto/bcrypt"
	"leaderboard/internal/database"
	"leaderboard/internal/models"
)

type userRepository struct{}

type UserRepository interface {
	FindByName(username string) (models.User, error)
	SaveToken(tokenValue string) (models.JwtToken, error)
	Store(u models.User) (models.User, error)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByName(username string) (models.User, error) {
	var user models.User
	db := database.DbManager()

	err := db.Where("username = ?", username).First(&user).Error

	return user, err
}

// Just save token for testing purposes and so it is available in database
func (r *userRepository) SaveToken(tokenValue string) (models.JwtToken, error) {
	var token models.JwtToken
	db := database.DbManager()
	token.Token = tokenValue

	err := db.Save(&token).Error

	return token, err
}

func (r *userRepository) Store(u models.User) (models.User, error) {
	db := database.DbManager()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)

	if err != nil {
		return u, err
	}

	u.Password = string(hashedPassword)

	err = db.Save(&u).Error

	return u, err
}
