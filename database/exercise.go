package database

import (
	"SC/config"
	"SC/models"
)

func CreateSetSoal(setSoal models.Set_soal) (models.Set_soal, error) {
	if err := config.DB.Save(&setSoal).Error; err != nil {
		return setSoal, err
	}
	return setSoal, nil
}

func LastQuestion() int {
	var soal models.Soal
	config.DB.Last(&soal)
	return int(soal.ID)
}

func GetOneSoal(id int) models.Soal {
	var soal models.Soal
	config.DB.Find(&soal, "id=?", id)
	return soal
}

func InputSetSoalDetail(setSoalDetail models.Set_soal_detail) {
	config.DB.Save(&setSoalDetail)
}
