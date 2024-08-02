package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/onsana/order_service/handlers"
	"github.com/onsana/order_service/service"
	"github.com/onsana/order_service/storage"
)

func SetupRoutes(app *fiber.App) {

	orderStorage := storage.NewOrderStorage()
	addressStorage := storage.NewAddressStorage()
	productStorage := storage.NewProductStorage()
	orderService := service.NewService(orderStorage, addressStorage, productStorage)
	orderHandler := handlers.NewHandler(orderService)

	app.Get("/orders", handlers.GetAllOrders)
	// app.Get("/orders/:id", handlers.GetOrderById)
	app.Post("/orders", orderHandler.CreateOrder)
	// app.Delete("/orders/:id", handlers.DeleteOrderById)
	// app.Put("/orders/:id", handlers.UpdateOrderDataById)
}
