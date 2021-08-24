package auth

import (
	"SC/models"
	"testing"

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
	userToken, err := CreateToken(1)
	if assert.NoError(t, err) {
		assert.NotEqual(t, "", userToken)
	}
}

func TestCreateTokenAdmin(t *testing.T) {
	userToken, err := CreateAdminToken(1)
	if assert.NoError(t, err) {
		assert.NotEqual(t, "", userToken)
	}
}

// func TestRegisterControllerSuccess(t *testing.T) {
// 	e := echo.New()
// 	token, _ := CreateToken(1)
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.Set("user", token)
// 	// fmt.Println(c.Get("user"))
// 	userId, role := ExtractTokenUserId(c)
// 	assert.Equal(t, 1, userId)
// 	assert.Equal(t, "user", role)
// 	// assert.Equal(t, http.StatusOK, rec.Code)
// 	// body := rec.Body.String()
// 	// var responseUser models.UserResponse
// 	// fmt.Println(body)
// 	// json.Unmarshal([]byte(body), &responseUser)

// 	// assert.Equal(t, true, responseUser.Status)
// 	// assert.Equal(t, "Registration success", responseUser.Message)

// }
