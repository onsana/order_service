package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/onsana/order_service/data"
	"github.com/onsana/order_service/database"
	"github.com/onsana/order_service/handlers"
	"github.com/onsana/order_service/service"
	"github.com/onsana/order_service/storage"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()
	app.Use(handlers.AuthMiddleware)
	setup(app)
	err := app.Listen(":6000")
	if err != nil {
		return
	}
}

func setup(app *fiber.App) {
	db := database.NewDBConnection()

	orderStorage := storage.NewOrderStorage(db)
	addressStorage := storage.NewAddressStorage(db)
	productStorage := storage.NewProductStorage(db)

	idToProductDto := data.CreateProductMock()
	productGateway := service.NewProductGatewayMock(idToProductDto)
	addressService := service.NewAddressService(addressStorage)
	productService := service.NewProductService(productStorage, productGateway)
	orderService := service.NewOrderService(orderStorage, *addressService, *productService)

	orderHandler := handlers.NewHandler(orderService)

	app.Get("/orders", orderHandler.GetAllOrders)
	app.Get("/orders/:id", orderHandler.GetOrderById)
	app.Post("/orders", orderHandler.CreateOrder)
	app.Put("/orders/:id", orderHandler.UpdateOrder)
	app.Delete("/orders/:id", orderHandler.DeleteOrderById)

	// app.Put("/orders/:id", handlers.UpdateOrderDataById)
}
