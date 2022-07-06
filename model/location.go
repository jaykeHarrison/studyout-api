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

func AddLocationFeatures(locationFeature *models.LocationFeature) error {
	return database.Database.Db.Create(&locationFeature).Error
}

func FetchLocationById(location *models.Location, locationId int64) error {
	return database.Database.Db.Joins("Users").First(&location, locationId).Error
}

func FetchFeaturesByLocationId(locationFeature *models.LocationFeature, locationId int64) error {
	return database.Database.Db.Joins("Location").First(&locationFeature, locationId).Error
}

func RemoveLocationById(location *models.Location, locationId int64) error {
	return database.Database.Db.Delete(&location, locationId).Error
}
