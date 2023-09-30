package main

import (
	"log"

	"github.com/Jagadish28/go-twitter/pkg/database"
	"github.com/Jagadish28/go-twitter/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	database.Migrate()
	database.Connect()

	routes.PublicRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
