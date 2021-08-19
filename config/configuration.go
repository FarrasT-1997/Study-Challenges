package config

import (
	"SC/models"
	"os"
	"strconv"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB
var HTTP_PORT int

func Init_DB() {
	connectionString := os.Getenv("CONNECTION_STRING")
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	var err error
	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Set_soal{})
	DB.AutoMigrate(&models.Set_soal_detail{})
	DB.AutoMigrate(&models.Soal{})
	DB.AutoMigrate(&models.Category{})
}
