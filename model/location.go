package model

import (
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/utils"
)

func FetchLocations(locationSlice *[]utils.Location) {
	database.Database.Db.Find(&locationSlice)
}
