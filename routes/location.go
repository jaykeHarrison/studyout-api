package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/utils"
)

func GetLocations(c *fiber.Ctx) error {
	//create empty slice whose elements are Location models
	locations := []utils.Location{}

	model.FetchLocations(&locations)

	return c.Status(200).JSON(locations)
}
