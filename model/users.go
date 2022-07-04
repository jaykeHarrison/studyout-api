package model

import (
	// "errors"
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/models"
)

func AddUser(user *models.Users) error {
	return database.Database.Db.Create(&user).Error
}

func FetchUserById(user *models.Users, userId int64) error {
	return database.Database.Db.First(&user, userId).Error
}

func RemoveUserById(User *models.Users, UserId int64) error {
	return database.Database.Db.Delete(&User, UserId).Error
}
