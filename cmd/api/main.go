package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"leaderboard/internal/controllers"
	"leaderboard/internal/database"
	"leaderboard/internal/handler"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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
	authController := controllers.NewAuthController()
	homeController := controllers.NewHomeController()
	scoreController := controllers.NewScoreController()
	leaderboardController := controllers.NewLeaderboardController()

	// API groups
	apiGroup := e.Group("/api/v1")

	publicGroup := apiGroup.Group("/")

	e.GET("", homeController.Home)
	publicGroup.POST("login", authController.Login)

	scoreGroup := apiGroup.Group("/scores")
	leaderboardGroup := apiGroup.Group("/leaderboard")

	authenticatedMiddleware := middleware.JWTWithConfig(getJWTConfig())

	scoreGroup.Use(authenticatedMiddleware)
	leaderboardGroup.Use(authenticatedMiddleware)

	// API routes
	scoreGroup.POST("", scoreController.Store)
	leaderboardGroup.GET("", leaderboardController.Index)
	leaderboardGroup.GET("/:type", leaderboardController.Index)

	e.Logger.Fatal(e.Start(":8080"))
}

func getJWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		SigningKey: []byte(viper.GetString("JWT_SECRET")),
	}
}
