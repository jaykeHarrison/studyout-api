package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/models"
	"github.com/jaykeHarrison/studyout-api/utils"
	"strconv"
	"reflect"
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

	// responseLocation := utils.CreateResponseLocation(fetchedLocation)

	return c.Status(200).JSON(bookmarks)

	// model.FetchBookmarks(&bookmarks,u)

    // if err:= model.FetchBookmarks(&bookmarks,u); err != nil {

	// 	return c.Status(400).SendString("400 : Bad request")
	// }

	// if len(bookmarks) == 0 {
	// 	return c.Status(404).SendString("404 : Not found")
	// }

	// return c.Status(200).JSON(bookmarks)

}