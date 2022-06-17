package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"leaderboard/internal/database"
	"leaderboard/internal/models"
	"leaderboard/internal/repositories"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	initRandom()
	configFileName := ".env"

	// Load Viper
	viper.SetConfigFile(configFileName)
	err := viper.ReadInConfig()

	if err != nil {
		panic("Unable to load configuration file: " + configFileName + ":" + err.Error())
	}

	// Load database
	err = database.Init()

	if err != nil {
		panic("Unable to establish database connection: " + err.Error())
	}

	scoreSeedCount := 10000

	args := os.Args

	if len(args) > 1 {
		scoreSeedCount, err = strconv.Atoi(args[1])

		if err != nil {
			panic("Failed to read score seed count argument")
		}
	}

	if err := SeedUsers(); err != nil {
		panic("Failed to seed uders")
	}

	if err := SeedScores(scoreSeedCount); err != nil {
		panic("Failed to seed scores")
	}

	fmt.Println("Succesfuly seeded database")
}

func SeedUsers() error {
	userRepository := repositories.NewUserRepository()

	firstUser := models.User{
		Username: "leader",
		Password: "leader",
	}

	secondUser := models.User{
		Username: "scorer",
		Password: "scorer",
	}

	if _, err := userRepository.Store(firstUser); err != nil {
		return err
	}

	if _, err := userRepository.Store(secondUser); err != nil {
		return err
	}

	return nil
}

func seedTokens() error {
	userRepository := repositories.NewUserRepository()

	for i := 0; i < 10; i++ {
		claims := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		}
		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))

		if err != nil {
			return err
		}

		_, err = userRepository.SaveToken(t)

		if err != nil {
			return err
		}

	}

	return nil
}

func SeedScores(scoreCount int) error {
	scoreRepository := repositories.NewScoreRepository()

	for i := 0; i < scoreCount; i++ {
		score := models.PlayerScore{
			Name:        RandStringRunes(16),
			Score:       rand.Intn(5000),
			SubmittedAt: randomDate(),
		}

		_, err := scoreRepository.Save(score)

		if err != nil {
			log.Println("Failed to seed score", score)
		}
	}

	return nil
}

func initRandom() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randomDate() time.Time {
	min := time.Date(2022, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2022, 6, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
