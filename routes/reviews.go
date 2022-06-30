package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/utils"
)

func GetReviews(c *fiber.Ctx) error {
	//create empty slice whose elements are Review models
	reviews := []utils.Review{}

	model.FetchReviews(&reviews)

	return c.Status(200).JSON(reviews)
}