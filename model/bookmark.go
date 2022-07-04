package model

import (
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/models"
	"github.com/jaykeHarrison/studyout-api/utils"
)


func AddBookmark(bookmark *models.Bookmark) error {
	return database.Database.Db.Create(&bookmark).Error
}

func FetchBookmarks(bookmarkSlice *[]utils.Bookmark, user_id uint64) error {
	return database.Database.Db.Where("user_id = ?", user_id).Find(&bookmarkSlice).Error
}

func RemoveBookmark(bookmark *models.Bookmark) error {
	return database.Database.Db.Where("user_id = ? AND location_id = ?",  bookmark.UserId, bookmark.LocationId).Delete(&bookmark).Error}
