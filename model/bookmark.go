package model

import (
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/models"
	"github.com/jaykeHarrison/studyout-api/utils"
)


func AddBookmark(bookmark *models.Bookmark) error {
	return database.Database.Db.Create(&bookmark).Error
}

func FetchBookmarks(bookmarkSlice *[]utils.Bookmark, user uint64) error {
	return database.Database.Db.Where("user = ?", user).Find(&bookmarkSlice).Error
}