package middlewares

import (
	"SC/constant"
	"fmt"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = int(userId)
	claims["role"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}

func ExtractTokenUserId(c echo.Context) (int, string) {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		userId := int(claims["userId"].(float64))
		role := fmt.Sprintf("%v", claims["role"])
		return userId, role
	}
	return 0, "a"
}
