package main

import (
	"github.com/spf13/viper"
	"leaderboard/internal/controllers"
	"leaderboard/internal/database"
	"leaderboard/internal/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = handler.NewValidator()

	configFileName := ".env"

	// Load Viper
	viper.SetConfigFile(configFileName)
	err := viper.ReadInConfig()

	if err != nil {
		e.Logger.Fatal("Unable to load configuration file " + configFileName + ":" + err.Error())
	}

	// Load database
	err = database.Init()

	if err != nil {
		e.Logger.Fatal("Unable to establish database connection: " + err.Error())
	}

	// Controllers
	homeController := controllers.NewHomeController()
	scoreController := controllers.NewScoreController()
	leaderboardController := controllers.NewLeaderboardController()

	// API groups
	apiGroup := e.Group("/api/v1")
	scoreGroup := apiGroup.Group("/scores")
	leaderboardGroup := apiGroup.Group("/leaderboard")

	// API routes
	apiGroup.GET("", homeController.Index)
	scoreGroup.POST("", scoreController.Store)
	leaderboardGroup.GET("", leaderboardController.Index)

	e.Logger.Fatal(e.Start(":8080"))
}
