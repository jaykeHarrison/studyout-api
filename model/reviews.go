package model

import (
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/models"
	"github.com/jaykeHarrison/studyout-api/utils"
	"time"
)

func AddReview(review *models.Review) error {
	database.Database.Db.Create(review)

	return database.Database.Db.Model(&review).Update("CreatedAt", time.Now()).Error
}

func FetchReviews(reviewSlice *[]utils.Review, location_refer string) error {
  return database.Database.Db.Where("location_refer = ?", location_refer).Find(&reviewSlice).Error
}