package model

import (
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/models"
)


func AddBookmark(bookmark *models.Bookmark) error {
	return database.Database.Db.Create(&bookmark).Error
}

func FetchBookmarks(bookmarkSlice *[]models.Bookmark) {
	database.Database.Db.Find(&bookmarkSlice)
}
