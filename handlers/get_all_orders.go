package handlers

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/onsana/order_service/database"
	"github.com/onsana/order_service/dto"
)

func GetAllOrders(c fiber.Ctx) error {
	var orders []dto.OrderDto

	result := database.DB.Db.Find(&orders)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("Помилка під час отримання замовлень")
	}

	return c.JSON(orders)
}
