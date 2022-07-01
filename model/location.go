package model

import (
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/models"
)

func FetchLocations(locationSlice *[]models.Location) {
	database.Database.Db.Find(&locationSlice)
}

func AddLocation(location *models.Location) error {
	return database.Database.Db.Create(&location).Error
}

func FetchLocationById(location *models.Location, locationId int64) error {
	return database.Database.Db.First(&location, locationId).Error
}
