package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/models"
	"github.com/jaykeHarrison/studyout-api/utils"
)

func PostBookmark (c *fiber.Ctx) error {

    var bookmark models.Bookmark

    if err := c.BodyParser(&bookmark); err != nil {
        return c.Status(400).JSON(err.Error())
    }

    if err := model.AddBookmark(&bookmark); err != nil {
        return c.Status(400).JSON(err.Error())
    }

    responseBookmark := utils.CreateResponseBookmark(bookmark)
    
    return c.Status(200).JSON(responseBookmark)

}


func GetBookmarks(c *fiber.Ctx) error {

	bookmarks := []models.Bookmark{}

	model.FetchBookmarks(&bookmarks)

    responseBookmark := utils.CreateResponseBookmark(bookmarks)

	return c.Status(200).JSON(responseBookmark)

}