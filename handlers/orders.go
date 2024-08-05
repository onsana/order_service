package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/onsana/order_service/data/dto"
	"github.com/onsana/order_service/data/model"
)

type OrderService interface {
	CreateOrder(orderDto *dto.OrderDto) (uuid.UUID, error)
	DeleteOrderById(id uuid.UUID) error
	GetAllOrders() model.Order
	GetOrderById(id uuid.UUID) (*model.Order, error)
	UpdateOrder(orderDto *dto.OrderDto) (*dto.OrderDto, error)
}

type AddressService interface {
	CreateAddress(addressDto *dto.Address, order model.Order) (*dto.Address, error)
}
type ProductService interface {
	CreateProducts(productsDto *[]dto.Product, order model.Order) (*[]dto.Product, error)
	ValidateProducts(productsDto *[]dto.Product) (*[]dto.Product, error)
}

type ProductGateway interface {
	GetExistingProducts(productsDto *[]dto.Product) (*[]dto.Product, []uuid.UUID)
}

type ProductGatewayImpl struct {
}

type ProductGatewayMock struct {
	IdToProductDto map[uuid.UUID]dto.Product
}

func (p *ProductGatewayImpl) GetExistingProducts(_ *[]dto.Product) (*[]dto.Product, []uuid.UUID) {
	// here should be invocation of Product Service instance
	products := make([]dto.Product, 2)
	//products = append(products, dto.Product{})
	return &products, nil
}

func (p *ProductGatewayMock) GetExistingProducts(productsDto *[]dto.Product) (*[]dto.Product, []uuid.UUID) {
	var absentIds []uuid.UUID

	for i := range *productsDto {
		id := (*productsDto)[i].ProductID
		product, ok := p.IdToProductDto[id]
		if !ok {
			absentIds = append(absentIds, id)
		} else {
			(*productsDto)[i].ProductName = product.ProductName
			(*productsDto)[i].Price = product.Price
		}
	}
	return productsDto, absentIds
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
	orderId, err := oH.oS.CreateOrder(orderDto)
	if err != nil {
		return c.Status(422).JSON(any(err.Error()))
	}
	return c.Status(201).JSON(any(fmt.Sprintf("Order created with id = %s", orderId)))
}

func (oH *OrderHandler) GetAllOrders(c fiber.Ctx) error {
	orders := oH.oS.GetAllOrders()
	fmt.Println("Orders:", orders)
	return c.Status(200).JSON(orders)
}

func (oH *OrderHandler) GetOrderById(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	order, err := oH.oS.GetOrderById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Order not found"})
	}

	return c.Status(200).JSON(order)
}
func (oH *OrderHandler) DeleteOrderById(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid ID format"})
	}

	err = oH.oS.DeleteOrderById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "order not found or could not be deleted"})
	}

	return c.Status(204).JSON(nil)
}

func (oH *OrderHandler) UpdateOrder(c fiber.Ctx) error {
	orderDto := new(dto.OrderDto)
	if err := c.Bind().JSON(orderDto); err != nil {
		return err
	}

	err := updateOrderDtoWithId(orderDto, c)
	if err != nil {
		return c.Status(404).JSON(any(err.Error()))
	}

	order, err := oH.oS.UpdateOrder(orderDto)
	if err != nil {
		return c.Status(404).JSON(any(err.Error()))
	}
	return c.Status(200).JSON(order)
}

func updateOrderDtoWithId(orderDto *dto.OrderDto, c fiber.Ctx) error {
	id := c.Params("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("Order cannot be update, incorrect uuid format")
	}
	orderDto.ID = parsedId
	return nil
}
