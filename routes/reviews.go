package routes

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/models"
	"github.com/jaykeHarrison/studyout-api/utils"
)

func GetReviewsByLocationID(c *fiber.Ctx) error {
	//Assign the query parameter for location ID to location_refer variable
	location_refer := c.Params("location_id")

	//create empty slice whose elements are Review models
	reviews := []utils.Review{}

	model.FetchReviews(&reviews, location_refer)

	if err := model.FetchReviews(&reviews, location_refer); err != nil {

		return c.Status(400).SendString("400 : Bad request")
	}

	if len(reviews) == 0 {
		return c.Status(404).SendString("404 : Not found")
	}

	return c.Status(200).JSON(reviews)
}

func PostReview(c *fiber.Ctx) error {
	var newReview models.Review

	if err := c.BodyParser(&newReview); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := model.AddReview(&newReview); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseReview := utils.CreateResponseReview(newReview)

	return c.Status(201).JSON(responseReview)
}
