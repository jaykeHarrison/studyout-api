package model

import (
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/utils"
)

func FetchReviews(reviewSlice *[]utils.Review, location_refer string) error {
	
return database.Database.Db.Where("location_refer = ?", location_refer).Find(&reviewSlice).Error

}