package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/routes"
	"log"
)

//handler function that returns an error
func welcome(c *fiber.Ctx) error {
	//c points to the fiber Context/Ctx struct
	//Ctx struct holds the HTTP request and response

	//sends a JSON string of "welcome" in the reponse.
	return c.SendString("Welcome")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Get("/api/locations", routes.GetLocations)
	app.Get("/api/reviews" , routes.GetReviews)
}

func main() {
	//connect to database
	database.ConnectDb()
	//initialise a Fiber app
	app := fiber.New()
	setUpRoutes(app)

	//Listen on PORT 3000, log error if it occurs
	log.Fatal(app.Listen(":3000"))
}
