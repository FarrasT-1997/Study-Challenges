package models

import "gorm.io/gorm"

type Soal struct {
	gorm.Model
	Soal       string `json:"soal" form:"soal" gorm:"not null"`
	PilihanA   string `json:"pilihan-a" form:"pilihan-a"`
	PilihanB   string `json:"pilihan-b" form:"pilihan-b"`
	PilihanC   string `json:"pilihan-c" form:"pilihan-c"`
	PilihanD   string `json:"pilihan-d" form:"pilihan-d"`
	Jawaban    string `json:"jawaban" form:"jawaban"`
	Kesulitan  string `json:"kesulitan" form:"kesulitan"`
	Solusi     string `json:"solusi" form:"solusi"`
	Approval   string `json:"approval" form:"approval"`
	KategoriID uint
}
