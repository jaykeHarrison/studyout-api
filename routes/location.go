package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/models"
	"github.com/jaykeHarrison/studyout-api/utils"
	"reflect"
	"strconv"
)

func GetLocations(c *fiber.Ctx) error {
	//create empty slice whose elements are Location models
	locations := []models.Location{}

	model.FetchLocations(&locations)

	return c.Status(200).JSON(locations)
}

func PostLocation(c *fiber.Ctx) error {
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

func GetLocationById(c *fiber.Ctx) error {
	var fetchedLocation models.Location

	locationIdString := c.Params("location_id")

	locationIdInt, _ := strconv.ParseInt(locationIdString, 10, 0)

	if reflect.TypeOf(locationIdInt).String() != "int64" {
		return c.Status(400).JSON("location_id must be an integer")
	}

	if locationIdInt < 1 {
		return c.Status(400).JSON("invalid location_id")
	}

	if err := model.FetchLocationById(&fetchedLocation, locationIdInt); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	responseLocation := utils.CreateResponseLocation(fetchedLocation)

	return c.Status(200).JSON(responseLocation)
}
