package controller

import (
	"SC/database"
	"SC/models"
	"net/http"

	"github.com/labstack/echo"
)

func SubmitQuestionAdmin(c echo.Context) error {
	submitSoal := models.Soal{}
	submitSoal.Approval = "sudah"
	c.Bind(&submitSoal)
	soal, err := database.CreateQuestion(submitSoal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	mapSoal := map[string]interface{}{
		"Soal":      soal.ID,
		"Pilihan A": user.Nama,
		"Pilihan B": user.Email,
		"Pilihan C": user.TotalPoin,
		"Pilihan D": user.TotalPoin,
		"Jawaban":   user.TotalPoin,
		"Kesulitan": user.TotalPoin,
		"Solusi":    user.TotalPoin,
		"Approval":  user.Rank,
		"Kategori":  user.Rank,
	}
}

func SubmitQuestion(c echo.Context) error {

}
