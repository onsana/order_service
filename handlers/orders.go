package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/onsana/order_service/data/dto"
	"github.com/onsana/order_service/data/model"
)

func GetAllOrders(c fiber.Ctx) error {
	var orders []model.Order
	// TODO temp changes
	//result := database.DB.Db.Find(&orders)
	//if result.Error != nil {
	//	log.Println(result.Error)
	//	return c.Status(fiber.StatusInternalServerError).SendString("Помилка під час отримання замовлень")
	//}

	return c.JSON(orders)
}

type OrderService interface {
	CreateOrder(orderDto *dto.OrderDto) (uuid.UUID, error)
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
