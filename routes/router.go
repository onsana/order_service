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

	addressService := service.NewAddressService(addressStorage)
	productService := service.NewProductService(productStorage)
	orderService := service.NewOrderService(orderStorage, *addressService, *productService)

	orderHandler := handlers.NewHandler(orderService)

	app.Get("/orders", orderHandler.GetAllOrders)
	// app.Get("/orders/:id", handlers.GetOrderById)
	app.Post("/orders", orderHandler.CreateOrder)
	// app.Delete("/orders/:id", handlers.DeleteOrderById)
	// app.Put("/orders/:id", handlers.UpdateOrderDataById)
}
