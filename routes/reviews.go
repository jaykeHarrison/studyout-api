package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/utils"
)

func GetReviewsByLocationID(c *fiber.Ctx) error {
	//Assign the query parameter for location ID to location_refer variable
	location_refer := c.Params("location_id")

	//create empty slice whose elements are Review models
	reviews := []utils.Review{}

	model.FetchReviews(&reviews, location_refer)

	if err:= model.FetchReviews(&reviews, location_refer); err != nil {


		return c.Status(400).SendString("400 : Bad request")
	}

	if len(reviews) == 0 {
		return c.Status(404).SendString("404 : Not found")
	}

	return c.Status(200).JSON(reviews)
}