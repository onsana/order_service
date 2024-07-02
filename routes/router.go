package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/onsana/order_service/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/orders", handlers.GetAllOrders)
	// app.Get("/orders/:id", handlers.GetOrderById)
	// app.Post("/orders", handlers.CreateNewOrder)
	// app.Delete("/orders/:id", handlers.DeleteOrderById)
	// app.Put("/orders/:id", handlers.UpdateOrderDataById)
}
