package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/models"
	"github.com/jaykeHarrison/studyout-api/utils"
	"reflect"
	"strconv"
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

func GetUserById(c *fiber.Ctx) error {
	var fetchedUser models.Users

	userIdString := c.Params("user_id")

	userIdInt, _ := strconv.ParseInt(userIdString, 10, 0)

	if reflect.TypeOf(userIdInt).String() != "int64" {
		return c.Status(400).JSON("user_id must be an integer")
	}

	if userIdInt < 1 {
		return c.Status(400).JSON("invalid user_id")
	}

	if err := model.FetchUserById(&fetchedUser, userIdInt); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	responseUser := utils.CreateResponseUser(fetchedUser)

	return c.Status(200).JSON(responseUser)
}
