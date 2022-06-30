package model

import (
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/utils"
	"github.com/jaykeHarrison/studyout-api/models"
)

func FetchLocations(locationSlice *[]utils.Location) {
	database.Database.Db.Find(&locationSlice)
}

func AddLocation(location *models.Location) error {
	return database.Database.Db.Create(&location).Error
}

