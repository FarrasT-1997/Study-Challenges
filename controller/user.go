package controller

import (
	"SC/auth"
	"SC/database"
	"SC/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func UserSignup(c echo.Context) error {
	addUser := models.User{}
	addUser.TotalPoin = 0
	addUser.Rank = "bronze"
	addUser.Role = "user"
	c.Bind(&addUser)
	user, err := database.CreateUser(addUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapUser := map[string]interface{}{
		"ID":         user.ID,
		"Name":       user.Nama,
		"Email":      user.Email,
		"Total Poin": user.TotalPoin,
		"Rank":       user.Rank,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "new user added",
		"data":    mapUser,
	})
}

func UserLogin(c echo.Context) error {
	userData := models.User{}
	c.Bind(&userData)

	user, err := database.LoginUsers(userData.Email, userData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "username or password is not correct",
		})
	}
	mapUserLogin := map[string]interface{}{
		"ID":    user.ID,
		"Name":  user.Nama,
		"Token": user.Token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Welcome",
		"users":   mapUserLogin,
	})
}

// AUTHORIZATION USER
func UserAuthorize(userId int, c echo.Context) error {
	userAuth, err := database.GetOneUser(userId)
	loggedInUserId, role := auth.ExtractTokenUserId(c)
	if loggedInUserId != userId || userAuth.Role != role || err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	return nil
}

func ShowUserProfile(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if err = UserAuthorize(userId, c); err != nil {
		return err
	}

	user, err := database.GetDetailUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot find the user",
		})
	}
	mapUser := map[string]interface{}{
		"ID":         user.ID,
		"Name":       user.Nama,
		"Email":      user.Email,
		"Total Poin": user.TotalPoin,
		"Rank":       user.Rank,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    mapUser,
	})
}

func UserLogout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, err := database.GetOneUser(id)
	logout.Token = ""
	user, err := database.EditUser(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "GOODBYE!",
		"data":    user.Nama,
	})
}

func EditUserProfile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if err = UserAuthorize(id, c); err != nil {
		return err
	}

	editUser, err := database.GetOneUser(id)
	c.Bind(&editUser)
	user, err := database.EditUser(editUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit data",
		})
	}
	mapUser := map[string]interface{}{
		"ID":         user.ID,
		"Name":       user.Nama,
		"Email":      user.Email,
		"Total Poin": user.TotalPoin,
		"Rank":       user.Rank,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Profile Updated!",
		"data":    mapUser,
	})
}
