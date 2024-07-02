package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/onsana/order_service/model"
)

func GetAllOrders(c fiber.Ctx) error {
	var orders []model.Order

	return c.JSON(orders)
}
