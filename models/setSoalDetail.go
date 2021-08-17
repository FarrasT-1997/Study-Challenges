package models

import "gorm.io/gorm"

type SetSoalDetail struct {
	gorm.Model
	SetSoalID uint
	Soal      []Soal
}
