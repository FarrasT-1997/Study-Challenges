package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama       string `json:"nama" form:"nama" gorm:"size:50;not null"`
	Email      string `json:"email" form:"email" gorm:"size:50;not null"`
	Password   string `json:"password" form:"password" gorm:"size:50;not null"`
	Poin       int    `json:"poin" form:"poin"`
	Rank       string `json:"rank" form:"rank" gorm:"size:50"`
	Peringkat  int    `json:"peringkat" form:"peringkat"`
	TotalSoal  int    `json:"total-soal" form:"total-soal"`
	TotalBenar int    `json:"total-benar" form:"total-benar"`
	TotalSalah int    `json:"total-salah" form:"total-salah"`
	Token      string `json:"token" form:"token"`
}
