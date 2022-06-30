package model

import (
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/models"
)

func AddUser(user *models.Users) {
	database.Database.Db.Create(&user)
}
