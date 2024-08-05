package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/onsana/order_service/data"
	"github.com/onsana/order_service/data/dto"
	"github.com/onsana/order_service/data/model"
	"github.com/onsana/order_service/handlers"
)

type orderStorage interface {
	CreateOrder(order *model.Order) error
	GetAllOrders() []model.Order
	GetOrderById(id uuid.UUID) (model.Order, error)
	DeleteOrderById(id uuid.UUID) error
}

type addressStorage interface {
	CreateAddress(address *model.Address) error
}

type productStorage interface {
	CreateProducts(order *[]model.Product) error
}

type orderService struct {
	oSt orderStorage
	aS  addressService
	pS  productService
}
type addressService struct {
	aSt addressStorage
}

type productService struct {
	pSt productStorage
	pG  handlers.ProductGateway
}

func NewOrderService(oSt orderStorage, aS addressService, pS productService) *orderService {
	return &orderService{
		oSt: oSt,
		aS:  aS,
		pS:  pS,
	}
}
func NewAddressService(aSt addressStorage) *addressService {
	return &addressService{aSt: aSt}
}

func NewProductService(pSt productStorage, pG handlers.ProductGateway) *productService {
	return &productService{pSt: pSt, pG: pG}
}

func NewProductGatewayMock(idToProductDto map[uuid.UUID]dto.Product) *handlers.ProductGatewayMock {
	return &handlers.ProductGatewayMock{
		IdToProductDto: idToProductDto,
	}
}

func NewProductGatewayImpl() *handlers.ProductGatewayImpl {
	return &handlers.ProductGatewayImpl{}
}

func (a *addressService) CreateAddress(addressDto *dto.Address, order model.Order) (*dto.Address, error) {
	address := data.ConvertAddress(*addressDto, order)
	err := a.aSt.CreateAddress(address)
	if err != nil {
		return addressDto, err
	}
	return data.ConvertAddressToDto(*address), nil
}

func (p *productService) CreateProducts(productsDto *[]dto.Product, order model.Order) (*[]dto.Product, error) {
	products := data.ConvertProduct(*productsDto, order)
	err := p.pSt.CreateProducts(products)
	if err != nil {
		return nil, err
	}
	return data.ConvertProductToDto(*products), nil
}

func (p *productService) ValidateProducts(productsDto *[]dto.Product) (*[]dto.Product, error) {
	realProducts, absentIds := p.pG.GetExistingProducts(productsDto)
	if len(absentIds) > 0 {
		return nil, fmt.Errorf("Order cannot be created due to the absence of products: %v !", absentIds)
	}
	return realProducts, nil
}

func (o *orderService) CreateOrder(orderDto *dto.OrderDto) (uuid.UUID, error) {
	validatedProducts, err := o.pS.ValidateProducts(&orderDto.Products)
	if err != nil {
		return uuid.Nil, err
	}

	order := data.ConvertOrder(*orderDto)
	err = o.oSt.CreateOrder(order)
	if err != nil {
		return uuid.Nil, err
	}

	_, err = o.aS.CreateAddress(&orderDto.Address, *order)
	if err != nil {
		return uuid.Nil, err
	}

	_, err = o.pS.CreateProducts(validatedProducts, *order)
	if err != nil {
		return uuid.Nil, err
	}

	return order.ID, nil
}
func (o *orderService) GetAllOrders() []model.Order {
	return o.oSt.GetAllOrders()
}
func (o *orderService) GetOrderById(id uuid.UUID) (model.Order, error) {
	return o.oSt.GetOrderById(id)
}

func (o *orderService) DeleteOrderById(id uuid.UUID) error {
	return o.oSt.DeleteOrderById(id)
}
