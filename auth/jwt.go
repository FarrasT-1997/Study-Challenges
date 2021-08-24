package auth

import (
	"SC/constant"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LogMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
	}))
}

func CreateToken(userId int) (*jwt.Token, string) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = float64(userId)
	claims["role"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, _ := token.SignedString([]byte(constant.SECRET_JWT))
	return token, stringToken
}

func CreateAdminToken(userId int) *jwt.Token {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = float64(userId)
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func CreateSignedStringUser(userId int) (string, error) {
	token, _ := CreateToken(userId)
	return token.SignedString([]byte(constant.SECRET_JWT))
}

func CreateSignedStringAdmin(userId int) (string, error) {
	token := CreateAdminToken(userId)
	return token.SignedString([]byte(constant.SECRET_JWT))
}

func ExtractTokenUserId(c echo.Context) (int, string) {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		role := fmt.Sprintf("%v", claims["role"])
		return userId, role
	}
	return 0, "a"
}
