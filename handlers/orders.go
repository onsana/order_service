package handlers

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/onsana/order_service/data/dto"
	"github.com/onsana/order_service/data/model"

	"github.com/gofiber/fiber/v3"
)

type OrderService interface {
	CreateOrder(orderDto *dto.OrderDto) uuid.UUID
	GetAllOrders() []model.Order
}

type AddressService interface {
	CreateAddress(addressDto *dto.Address, order model.Order) *dto.Address
}
type ProductService interface {
	CreateProducts(productsDto *[]dto.Product, order model.Order) *[]dto.Product
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

func (oH *OrderHandler) GetAllOrders(c fiber.Ctx) error {
	orders := oH.oS.GetAllOrders()
	return c.Status(200).JSON(orders)
}
