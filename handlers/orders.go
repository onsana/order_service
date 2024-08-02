package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/onsana/order_service/data/dto"
	"github.com/onsana/order_service/data/model"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/onsana/order_service/database"
)

func GetAllOrders(c fiber.Ctx) error {
	var orders []model.Order

	result := database.DB.Db.Find(&orders)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("Помилка під час отримання замовлень")
	}

	return c.JSON(orders)
}

type OrderService interface {
	CreateOrder(orderDto *dto.OrderDto) uuid.UUID
}

type OrderHandler struct {
	oS OrderService
}

func NewHandler(s OrderService) OrderHandler {
	return OrderHandler{oS: s}
}

func (oH *OrderHandler) CreateOrder(c fiber.Ctx) error {
	orderDto := new(dto.OrderDto)

	if err := c.Bind().JSON(orderDto); err != nil {
		return err
	}
	orderId := oH.oS.CreateOrder(orderDto)

	return c.Status(200).JSON(any(fmt.Sprintf("Order created with id = %s", orderId)))
}
