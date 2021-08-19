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

func RandomId(soalCategory_id, level uint) []models.Soal {
	var soal []models.Soal
	config.DB.Raw("SELECT id FROM soals WHERE kesulitan_id = ? AND category_id = ? AND approval = 'accept' ORDER BY rand() LIMIT 5", level, soalCategory_id).Scan(&soal)
	return soal
}

func InputSetSoalDetail(setSoalDetail models.Set_soal_detail) {
	config.DB.Save(&setSoalDetail)
}

func ShowActiveSoal(setSoalId int) []models.Soal {
	var soalDetail []models.Set_soal_detail
	var soal []models.Soal
	config.DB.Raw("select soals.id, soal_pertanyaan, pilihan_a, pilihan_b, pilihan_c, pilihan_d from soals inner join set_soal_details on set_soal_details.soal_id = soals.id where set_soal_details.set_soal_id = ?", setSoalId).Scan(&soal).Scan(&soalDetail)
	return soal
}
