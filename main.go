package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/onsana/order_service/routes"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	routes.SetupRoutes(app)
	app.Listen(":6000")
}
