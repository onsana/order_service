package handlers

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/onsana/order_service/database"
	"github.com/onsana/order_service/dto"
)

func CreateNewOrder(c fiber.Ctx) error {
	order := new(dto.OrderDto)

	if err := c.Bind().JSON(order); err != nil {
		return err
	}
	log.Println(order)
	// database.DB.Db.Create(&order)

	database.DB.Db.Create(&dto.OrderDto{})

	return c.Status(200).JSON(order)
}
