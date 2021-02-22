package database

import (
	"project/config"
	"project/models"
)

func GetUsers() (interface{}, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUsers(user *models.Users) (interface{}, error) {

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
