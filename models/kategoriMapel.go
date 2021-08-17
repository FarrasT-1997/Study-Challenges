package models

import "gorm.io/gorm"

type Kategori struct {
	ID   uint   `json:"id" form:"id" gorm:"not null"`
	Nama string `json:"nama" form:"nama"`
}

type Matematika struct {
	gorm.Model
	Soal       string `json:"soal" form:"soal" gorm:"not null"`
	PilihanA   string `json:"pilihan-a" form:"pilihan-a"`
	PilihanB   string `json:"pilihan-b" form:"pilihan-b"`
	PilihanC   string `json:"pilihan-c" form:"pilihan-c"`
	PilihanD   string `json:"pilihan-d" form:"pilihan-d"`
	Jawaban    string `json:"jawaban" form:"jawaban"`
	Kesulitan  string `json:"kesulitan" form:"kesulitan"`
	Solusi     string `json:"solusi" form:"solusi"`
	KategoriID uint
}

type Fisika struct {
	gorm.Model
	Soal       string `json:"soal" form:"soal" gorm:"not null"`
	PilihanA   string `json:"pilihan-a" form:"pilihan-a"`
	PilihanB   string `json:"pilihan-b" form:"pilihan-b"`
	PilihanC   string `json:"pilihan-c" form:"pilihan-c"`
	PilihanD   string `json:"pilihan-d" form:"pilihan-d"`
	Jawaban    string `json:"jawaban" form:"jawaban"`
	Kesulitan  string `json:"kesulitan" form:"kesulitan"`
	Solusi     string `json:"solusi" form:"solusi"`
	KategoriID uint
}

type Kimia struct {
	gorm.Model
	Soal       string `json:"soal" form:"soal" gorm:"not null"`
	PilihanA   string `json:"pilihan-a" form:"pilihan-a"`
	PilihanB   string `json:"pilihan-b" form:"pilihan-b"`
	PilihanC   string `json:"pilihan-c" form:"pilihan-c"`
	PilihanD   string `json:"pilihan-d" form:"pilihan-d"`
	Jawaban    string `json:"jawaban" form:"jawaban"`
	Kesulitan  string `json:"kesulitan" form:"kesulitan"`
	Solusi     string `json:"solusi" form:"solusi"`
	KategoriID uint
}

type Biologi struct {
	gorm.Model
	Soal       string `json:"soal" form:"soal" gorm:"not null"`
	PilihanA   string `json:"pilihan-a" form:"pilihan-a"`
	PilihanB   string `json:"pilihan-b" form:"pilihan-b"`
	PilihanC   string `json:"pilihan-c" form:"pilihan-c"`
	PilihanD   string `json:"pilihan-d" form:"pilihan-d"`
	Jawaban    string `json:"jawaban" form:"jawaban"`
	Kesulitan  string `json:"kesulitan" form:"kesulitan"`
	Solusi     string `json:"solusi" form:"solusi"`
	KategoriID uint
}

type Ekonomi struct {
	gorm.Model
	Soal       string `json:"soal" form:"soal" gorm:"not null"`
	PilihanA   string `json:"pilihan-a" form:"pilihan-a"`
	PilihanB   string `json:"pilihan-b" form:"pilihan-b"`
	PilihanC   string `json:"pilihan-c" form:"pilihan-c"`
	PilihanD   string `json:"pilihan-d" form:"pilihan-d"`
	Jawaban    string `json:"jawaban" form:"jawaban"`
	Kesulitan  string `json:"kesulitan" form:"kesulitan"`
	Solusi     string `json:"solusi" form:"solusi"`
	KategoriID uint
}

type Geografi struct {
	gorm.Model
	Soal       string `json:"soal" form:"soal" gorm:"not null"`
	PilihanA   string `json:"pilihan-a" form:"pilihan-a"`
	PilihanB   string `json:"pilihan-b" form:"pilihan-b"`
	PilihanC   string `json:"pilihan-c" form:"pilihan-c"`
	PilihanD   string `json:"pilihan-d" form:"pilihan-d"`
	Jawaban    string `json:"jawaban" form:"jawaban"`
	Kesulitan  string `json:"kesulitan" form:"kesulitan"`
	Solusi     string `json:"solusi" form:"solusi"`
	KategoriID uint
}

type BahasaInggris struct {
	gorm.Model
	Soal       string `json:"soal" form:"soal" gorm:"not null"`
	PilihanA   string `json:"pilihan-a" form:"pilihan-a"`
	PilihanB   string `json:"pilihan-b" form:"pilihan-b"`
	PilihanC   string `json:"pilihan-c" form:"pilihan-c"`
	PilihanD   string `json:"pilihan-d" form:"pilihan-d"`
	Jawaban    string `json:"jawaban" form:"jawaban"`
	Kesulitan  string `json:"kesulitan" form:"kesulitan"`
	Solusi     string `json:"solusi" form:"solusi"`
	KategoriID uint
}
