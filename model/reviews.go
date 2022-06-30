package model

import (
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/utils"
)

func FetchReviews(reviewSlice *[]utils.Review) {
	database.Database.Db.Find(&reviewSlice)
}