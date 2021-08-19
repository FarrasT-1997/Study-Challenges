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

func PutAnswer(setSoalId int, jawabanUser map[int]string) {
	var soalDetail []models.Set_soal_detail
	config.DB.Where("set_soal_id=?", setSoalId).Find(&soalDetail)
	for i := 0; i < 5; i++ {
		soalDetail[i].Jawaban_user = jawabanUser[i+1]
		soalDetail[i].Status = "answered"
	}
	config.DB.Save(&soalDetail)
}

func Scoring(setSoalId int) (int, int, int, int, []int) {
	var soalDetail []models.Set_soal_detail
	var set_soal models.Set_soal
	var soal models.Soal

	var totalScore int = 0
	var totalBenar int = 0
	var totalSalah int = 0
	var totalTerjawab int = 0
	var SoalId_salah []int

	config.DB.Where("set_soal_id=?", setSoalId).Find(&soalDetail)
	config.DB.Where("id=?", setSoalId).Find(&set_soal)

	for i := 0; i < 5; i++ {
		config.DB.Where("id=?", soalDetail[i].SoalID).Find(&soal)
		if soal.Jawaban == soalDetail[i].Jawaban_user {
			totalBenar++
			totalTerjawab++
			switch set_soal.KesulitanID {
			case 1:
				soalDetail[i].Poin = 2
			case 2:
				soalDetail[i].Poin = 3
			case 3:
				soalDetail[i].Poin = 4
			}
		} else if soal.Jawaban != soalDetail[i].Jawaban_user && soalDetail[i].Jawaban_user == "pass" {
			soalDetail[i].Poin = 0
		} else if soal.Jawaban != soalDetail[i].Jawaban_user && soalDetail[i].Jawaban_user != "pass" {
			totalSalah++
			totalTerjawab++
			SoalId_salah = append(SoalId_salah, int(soal.ID))
			soalDetail[i].Poin = -1
		}
		totalScore += soalDetail[i].Poin
	}
	config.DB.Save(&soalDetail)
	return totalScore, totalBenar, totalSalah, totalTerjawab, SoalId_salah
}
