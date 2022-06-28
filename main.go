package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

//handler function that returns an error
func welcome(c *fiber.Ctx) error {
	//c points to the fiber Context/Ctx struct
	//Ctx struct holds the HTTP request and response
	return c.SendString("Welcome")
}

func main() {
	//initialise a Fiber app
	app := fiber.New()

	//define a route/endpoint with a handler function 'welcome'
	app.Get("/api", welcome)

	//Listen on PORT 3000, log error if it occurs
	log.Fatal(app.Listen(":3000"))
}
