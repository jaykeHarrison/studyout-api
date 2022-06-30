package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/models"
	"github.com/jaykeHarrison/studyout-api/utils"
)

func PostUser(c *fiber.Ctx) error {
	var newUser models.Users

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := model.AddUser(&newUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := utils.CreateResponseUser(newUser)

	return c.Status(200).JSON(responseUser)
}
