package database

import (
	"SC/config"
	"SC/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockSetSoal = models.Set_soal{
		KesulitanID:   1,
		TotalBenar:    0,
		TotalSalah:    0,
		TotalTerjawab: 0,
		Status:        "not answered yet",
		CategoryID:    2,
		UserID:        1,
	}
	mockSoal1 = models.Soal{
		Soal_pertanyaan: "1+1 = ?",
		PilihanA:        "2",
		PilihanB:        "3",
		PilihanC:        "4",
		PilihanD:        "1",
		Jawaban:         "a",
		KesulitanID:     1,
		Solusi:          "Ketika 1 ditambah 1 maka hasilnya pasti 2",
		Approval:        "accept",
		CategoryID:      4,
	}

	mockSoal2 = models.Soal{
		Soal_pertanyaan: "2+2 = ?",
		PilihanA:        "2",
		PilihanB:        "3",
		PilihanC:        "4",
		PilihanD:        "1",
		Jawaban:         "c",
		KesulitanID:     1,
		Solusi:          "Ketika 2 ditambah 2 maka hasilnya pasti 4",
		Approval:        "accept",
		CategoryID:      4,
	}

	mockSoal3 = models.Soal{
		Soal_pertanyaan: "Sebuah balok bermassa 1,5 kg didorong ke atas oleh gaya konstan F = 15 N pada bidang miring seperti gambar. Anggap percepatan gravitasi (g) 10 ms-2 dan gesekan antara balok dan bidang miring nol. Usaha total yang dilakukan pada balok adalah ...",
		PilihanA:        "15 J",
		PilihanB:        "30 J",
		PilihanC:        "35 J",
		PilihanD:        "45 J",
		Jawaban:         "a",
		KesulitanID:     1,
		Solusi:          "WTOT = ΣF . s = (15 – 7,5) . 2 = 15 joule",
		Approval:        "accept",
		CategoryID:      2,
	}
	mockSoal4 = models.Soal{
		Soal_pertanyaan: "Sebuah mobil mulai bergerak dari keadaan diam dengan percepatan tetap 24 m/s2. Maka kecepatan mobil setelah bergerak selama 18 sekon adalah:",
		PilihanA:        "2 m/s",
		PilihanB:        "24 m/s",
		PilihanC:        "36 m/s",
		PilihanD:        "42 m/s",
		Jawaban:         "d",
		KesulitanID:     1,
		Solusi:          "vt = v0 + at = 0 + (24 m/s2) (18 s). vt = 42 m/s",
		Approval:        "accept",
		CategoryID:      2,
	}
	mockSetSoalDetail = models.Set_soal_detail{
		Set_soalID:   1,
		SoalID:       3,
		Status:       "not answered yet",
		Poin:         0,
		Jawaban_user: "pass",
	}
	mockSetSoalDetail2 = models.Set_soal_detail{
		Set_soalID:   1,
		SoalID:       4,
		Status:       "not answered yet",
		Poin:         0,
		Jawaban_user: "pass",
	}
)

func TestCreateSetSoalSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Set_soal{})
	config.DB.Migrator().AutoMigrate(&models.Set_soal{})
	createdSetSoal, err := CreateSetSoal(mockSetSoal)
	if assert.NoError(t, err) {
		assert.Equal(t, uint(1), createdSetSoal.UserID)
		assert.Equal(t, uint(2), createdSetSoal.CategoryID)
		assert.Equal(t, uint(1), createdSetSoal.KesulitanID)
	}
}

func TestCreateSetSoalError(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Set_soal{})
	_, err := CreateSetSoal(mockSetSoal)
	assert.Error(t, err)
}

func TestGetRandomSoalSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Soal{})
	config.DB.Migrator().AutoMigrate(&models.Soal{})
	CreateQuestion(mockSoal1)
	CreateQuestion(mockSoal2)
	CreateQuestion(mockSoal3)
	CreateQuestion(mockSoal4)
	randomSoal := RandomId(uint(4), uint(1))
	if randomSoal[0].ID == 1 {
		assert.Equal(t, uint(1), randomSoal[0].ID)
	}
	if randomSoal[0].ID == 2 {
		assert.Equal(t, uint(2), randomSoal[0].ID)
	}
	if randomSoal[1].ID == 1 {
		assert.Equal(t, uint(1), randomSoal[1].ID)
	}
	if randomSoal[1].ID == 2 {
		assert.Equal(t, uint(2), randomSoal[1].ID)
	}
}

func TestInputSetSoalDetailSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Set_soal_detail{})
	config.DB.Migrator().AutoMigrate(&models.Set_soal_detail{})
	createdSetSoalDetail := InputSetSoalDetail(mockSetSoalDetail)
	assert.Equal(t, uint(1), createdSetSoalDetail.Set_soalID)
	assert.Equal(t, uint(3), createdSetSoalDetail.SoalID)
}

func TestShowSetSoalSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Set_soal{})
	config.DB.Migrator().AutoMigrate(&models.Set_soal{})
	createdSetSoal, _ := CreateSetSoal(mockSetSoal)
	setSoalShowed := ShowSetSoal(int(createdSetSoal.ID))
	assert.Equal(t, uint(1), setSoalShowed.UserID)
	assert.Equal(t, uint(2), setSoalShowed.CategoryID)
	assert.Equal(t, uint(1), setSoalShowed.KesulitanID)
}

func TestShowActiveSoalSuccess(t *testing.T) {
	config.Init_DB_Test()
	config.DB.Migrator().DropTable(&models.Set_soal_detail{})
	config.DB.Migrator().AutoMigrate(&models.Set_soal_detail{})
	config.DB.Migrator().DropTable(&models.Soal{})
	config.DB.Migrator().AutoMigrate(&models.Soal{})
	createdSetSoalDetail1 := InputSetSoalDetail(mockSetSoalDetail)
	createdSetSoalDetail2 := InputSetSoalDetail(mockSetSoalDetail2)
	CreateQuestion(mockSoal1)
	CreateQuestion(mockSoal2)
	CreateQuestion(mockSoal3)
	CreateQuestion(mockSoal4)

	activeSoalShowed := ShowActiveSoal(int(createdSetSoalDetail1.Set_soalID))
	assert.Equal(t, createdSetSoalDetail1.SoalID, activeSoalShowed[0].ID)
	assert.Equal(t, createdSetSoalDetail2.SoalID, activeSoalShowed[1].ID)
}
