package routes

import (
	"reflect"
	"strconv"

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

	bookmarks := []utils.Bookmark{}
	userIdString := c.Params("user_id")

	userIdInt, _ := strconv.ParseUint(userIdString, 10, 0)

	if reflect.TypeOf(userIdInt).String() != "uint64" {
		return c.Status(400).JSON("user_id must be an integer")
	}

	if userIdInt < 1 {
		return c.Status(400).JSON("invalid user_id")
	}

	if err := model.FetchBookmarks(&bookmarks, userIdInt); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(bookmarks)

}


func DeleteBookmark(c *fiber.Ctx)error{
	var bookmark models.Bookmark
	
	if err := c.BodyParser(&bookmark); err != nil {
        return c.Status(400).JSON(err.Error())
    }
	
	if err := model.RemoveBookmark(&bookmark); err !=nil {
		return c.Status(400).JSON(err.Error())
	}
	if bookmark.UserId < 1 {
		return c.Status(400).JSON("invalid user_id")
	}
	if bookmark.LocationId < 1 {
		return c.Status(400).JSON("invalid Location Id")
	}
	
	return c.Status(200).SendString("Bookmark has been deleted")


	

}
