package main

import (
	"fmt"
	"github.com/spf13/viper"
	"leaderboard/internal/database"

	"leaderboard/internal/models"
)

func main() {
	configFileName := ".env"

	// Load Viper
	viper.SetConfigFile(configFileName)
	err := viper.ReadInConfig()

	if err != nil {
		panic("Unable to load configuration file " + configFileName + ":" + err.Error())
	}

	// Load database
	err = database.Init()

	if err != nil {
		panic("Unable to establish database connection: " + err.Error())
	}

	err = Migrate()

	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	fmt.Println("Succesfuly migrated database")
}

func Migrate() error {
	db := database.DbManager()

	err := db.AutoMigrate(&models.PlayerScore{})

	if err != nil {
		return err
	}

	return nil
}
