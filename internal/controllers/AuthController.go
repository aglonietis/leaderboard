package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"leaderboard/internal/handler"
	"leaderboard/internal/repositories"
	"leaderboard/internal/requests"
	"net/http"
	"time"
)

type authController struct{
	userRepository repositories.UserRepository

}

type AuthController interface {
	Login(ctx echo.Context) error
}

func NewAuthController() AuthController {
	return &authController{
		userRepository: repositories.NewUserRepository(),
	}
}

func (c *authController) Login(ctx echo.Context) error {

	var request requests.LoginRequest

	err := handler.BindValidate(ctx, &request)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}


	user,err := c.userRepository.FindByName(request.Username)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, "")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password));

	if err != nil {
		return ctx.JSON(http.StatusNotFound, "")
	}

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

	jwtToken, err := c.userRepository.SaveToken(t)

	return ctx.JSON(http.StatusOK, jwtToken)
}
