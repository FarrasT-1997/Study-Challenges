package auth

import (
	"SC/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDBUser = models.User{
		Nama:      "farras",
		Email:     "farras@gmail.com",
		Password:  "123",
		TotalPoin: 0,
		Rank:      "bronze",
		Role:      "user",
	}
	mockDBUserLogin = models.User{
		Email:    "farras@gmail.com",
		Password: "123",
	}
)

func TestCreateTokenUser(t *testing.T) {
	userToken, err := CreateSignedStringUser(1)
	if assert.NoError(t, err) {
		assert.NotEqual(t, "", userToken)
	}
}

func TestCreateTokenAdmin(t *testing.T) {
	userToken, err := CreateSignedStringAdmin(1)
	if assert.NoError(t, err) {
		assert.NotEqual(t, "", userToken)
	}
}

func TestRegisterControllerSuccess(t *testing.T) {
	userId := 73
	token, signedString := CreateToken(userId)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", signedString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	token.Valid = true
	context.Set("user", token)
	userId, role := ExtractTokenUserId(context)
	assert.Equal(t, 73, userId)
	assert.Equal(t, "user", role)
}

func TestRegisterControllerNotValid(t *testing.T) {
	userId := 73
	token, signedString := CreateToken(userId)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", signedString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.Set("user", token)
	userId, role := ExtractTokenUserId(context)
	assert.Equal(t, 0, userId)
	assert.Equal(t, "a", role)

}
