package models

import "gorm.io/gorm"

type SetSoal struct {
	gorm.Model
	JawabanSatu   string `json:"jawaban-satu" form:"jawaban-satu"`
	JawabanDua    string `json:"jawaban-dua" form:"jawaban-dua"`
	JawabanTiga   string `json:"jawaban-tiga" form:"jawaban-tiga"`
	JawabanEmpat  string `json:"jawaban-empat" form:"jawaban-empat"`
	JawabanLima   string `json:"jawaban-lima" form:"jawaban-lima"`
	KesulitanID   uint
	TotalBenar    int    `json:"total-benar" form:"total-benar"`
	TotalSalah    int    `json:"total-salah" form:"total-salah"`
	TotalTerjawab int    `json:"total-terjawab" form:"total-terjawab"`
	Status        string `json:"status" form:"status"`
	KategoriID    uint
}
