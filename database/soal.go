package database

import (
	"SC/config"
	"SC/models"
)

func SubmitQuestion(soal models.Soal) {
	if err := config.DB.Save(&soal).Error; err != nil {
		return
	}
}
