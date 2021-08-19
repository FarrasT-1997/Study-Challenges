package controller

import (
	"SC/database"
	"SC/models"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GenerateRandomQuestion(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if err = UserAuthorize(userId, c); err != nil {
		return err
	}

	setSoal := models.Set_soal{
		TotalBenar:    0,
		TotalSalah:    0,
		TotalTerjawab: 0,
		Status:        "Not Answered Yet",
		UserID:        uint(userId),
	}
	c.Bind(&setSoal)
	newSetSoal, err := database.CreateSetSoal(setSoal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}

	randomizeSoal(newSetSoal.ID, newSetSoal.CategoryID, newSetSoal.KesulitanID, c)

	materi, level := levelAndCategoryIDConvert(newSetSoal.CategoryID, newSetSoal.KesulitanID)
	mapSetSoal := map[string]interface{}{
		"ID":             newSetSoal.ID,
		"UserID":         newSetSoal.UserID,
		"Mata Pelajaran": materi,
		"Kesulitan":      level,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "new set soal added",
		"data":    mapSetSoal,
	})
}

func levelAndCategoryIDConvert(materi, level uint) (string, string) {
	var materiConvert string
	var levelConvert string

	switch materi {
	case 1:
		materiConvert = "Kimia"
	case 2:
		materiConvert = "Fisika"
	case 3:
		materiConvert = "Biologi"
	case 4:
		materiConvert = "Matematika"
	case 5:
		materiConvert = "Bahasa Inggris"
	case 6:
		materiConvert = "Ekonomi"
	case 7:
		materiConvert = "Geografi"
	default:
		materiConvert = "Unknown"
	}

	switch level {
	case 1:
		levelConvert = "Easy"
	case 2:
		levelConvert = "Medium"
	case 3:
		levelConvert = "Hard"
	default:
		levelConvert = "unknown"
	}
	return materiConvert, levelConvert
}

func randomizeSoal(soal_id, soalCategory_id, level uint, c echo.Context) {
	var recordedRandom = []int{}
	sameRandom := true
	same := false
	min := 1
	max := database.LastQuestion()
	var randomId int
	var selectedSoal models.Soal
	for i := 1; i <= 5; i++ {
		randomId = min + rand.Intn(max-min+1)
		fmt.Println(randomId)
		for j := 0; j < len(recordedRandom); j++ {
			if recordedRandom[j] == randomId {
				same = true
				j = len(recordedRandom)
			}
		}
		recordedRandom = append(recordedRandom, randomId)
		selectedSoal = database.GetOneSoal(randomId)
		if sameRandom == same || selectedSoal.Approval != "accept" || selectedSoal.CategoryID != soalCategory_id || selectedSoal.KesulitanID != level {
			i--
			continue
		}
		newSetSoalDetail := models.Set_soal_detail{
			Set_soalID:   soal_id,
			SoalID:       selectedSoal.ID,
			Status:       "not answered",
			Poin:         0,
			Jawaban_user: "pass",
		}
		database.InputSetSoalDetail(newSetSoalDetail)
	}
}
