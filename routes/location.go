package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/model"
    "github.com/jaykeHarrison/studyout-api/models"
	"github.com/jaykeHarrison/studyout-api/utils"
)

func GetLocations(c *fiber.Ctx) error {
	//create empty slice whose elements are Location models
	locations := []models.Location{}

	model.FetchLocations(&locations)

	return c.Status(200).JSON(locations)
}



func PostLocation (c *fiber.Ctx) error {
    var location models.Location

    if err := c.BodyParser(&location); err != nil {
        return c.Status(400).JSON(err.Error())
    }

    if err := model.AddLocation(&location); err != nil {
        return c.Status(400).JSON(err.Error())
    }

    responseLocation := utils.CreateResponseLocation(location)
    
    return c.Status(200).JSON(responseLocation)

}



