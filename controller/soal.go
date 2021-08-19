package controller

import (
	"SC/auth"
	"SC/database"
	"SC/models"
	"net/http"

	"github.com/labstack/echo"
)

func Authorized(c echo.Context) (bool, models.User) {
	userId, token := auth.ExtractTokenUserId(c)
	userList, _ := database.GetOneUser(userId)

	if userList.Token != token {
		return false, userList
	}
	return true, userList
}

func SubmitQuestion(c echo.Context) error {
	auth, userList := Authorized(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
}
