package database

import (
	"SC/config"
	"SC/models"
)

func GetOneUser(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}
