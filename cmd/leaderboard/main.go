package main

import (
	"github.com/spf13/viper"
	"leaderboard/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configFileName := ".env"

	viper.SetConfigFile(configFileName)
	err := viper.ReadInConfig()

	if err != nil {
		e.Logger.Fatal("Unable to load configuration file " + configFileName)
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
	scoreGroup.GET("", scoreController.Test)
	leaderboardGroup.GET("", leaderboardController.Test)

	e.Logger.Fatal(e.Start(":8080"))
}
