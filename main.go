package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/onsana/order_service/database"
	"github.com/onsana/order_service/middleware"
	"github.com/onsana/order_service/routes"
)

func main() {
	database.ConnectDb()
	// Initialize a new Fiber app
	app := fiber.New()
	app.Use(middleware.AuthMiddleware)
	routes.SetupRoutes(app)
	app.Listen(":6000")
}
