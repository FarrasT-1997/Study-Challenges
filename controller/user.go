package controller

import (
	"SC/auth"
	"SC/database"
	"SC/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func ourEncrypt(plain string) string {
	bytePlain := []byte(plain)
	hashed, _ := bcrypt.GenerateFromPassword(bytePlain, bcrypt.MinCost)
	return string(hashed)
}

func UserSignup(c echo.Context) error {
	input := models.User{}
	c.Bind(&input)
	if input.Nama == "" || input.Email == "" || input.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please fill name, email and password correctly",
		})
	}
	if same, _ := database.CheckSameEmail(input.Email); same == true {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "email already used",
		})
	}
	addUser := models.User{}
	addUser.Nama = input.Nama
	addUser.Email = input.Email
	addUser.Password = ourEncrypt(input.Password)
	addUser.TotalPoin = 0
	addUser.Rank = "bronze"
	addUser.Role = "user"
	fmt.Println(addUser.Password)
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
	input := models.User{}
	c.Bind(&input)
	userData := models.User{
		Nama:     input.Email,
		Password: ourEncrypt(input.Password),
	}
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
