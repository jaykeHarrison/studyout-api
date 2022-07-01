package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/studyout-api/database"
	"github.com/jaykeHarrison/studyout-api/database/seed"
	"github.com/jaykeHarrison/studyout-api/routes"
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
	app.Get("/api/reviews/:location_id", routes.GetReviewsByLocationID)
	app.Get("/api/bookmarks/:user_id", routes.GetBookmarks)
  app.Get("/api/users/:user_id", routes.GetUserById)
	app.Get("/api/locations/:location_id", routes.GetLocationById)
	app.Post("/api/users", routes.PostUser)
	app.Post("/api/reviews", routes.PostReview)
	app.Post("/api/locations", routes.PostLocation)
	app.Post("/api/bookmarks", routes.PostBookmark)
	app.Delete("/api/locations/:location_id", routes.DeleteLocationById)
}


func main() {
	//connect to database
	database.ConnectDb()
	//initialise a Fiber app
	app := fiber.New()
	setUpRoutes(app)
	seed.Seed()
	//Listen on PORT 3000, log error if it occurs
	log.Fatal(app.Listen(":3000"))
}
