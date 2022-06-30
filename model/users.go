package model

import (
	// "errors"
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/models"
)

func AddUser(user *models.Users) error {
	return database.Database.Db.Create(&user).Error
}
