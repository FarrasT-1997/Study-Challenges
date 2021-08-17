package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama" gorm:"size:50;not null"`
	Email    string `json:"email" form:"email" gorm:"size:50;not null"`
	Password string `json:"password" form:"password" gorm:"size:50;not null"`
	Token    string `json:"token" form:"token"`
}
