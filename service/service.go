package service

import (
	"github.com/google/uuid"
	"github.com/onsana/order_service/data"
	"github.com/onsana/order_service/data/dto"
	"github.com/onsana/order_service/data/model"
)

type orderStorage interface {
	CreateOrder(order *model.Order) model.Order
}

type addressStorage interface {
	CreateAddress(address *model.Address) model.Address
}

type productStorage interface {
	CreateProducts(order *[]model.Product) []model.Product
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
func NewProductService(pSt productStorage) *productService {
	return &productService{pSt: pSt}
}

func (a *addressService) CreateAddress(addressDto *dto.Address, order model.Order) *dto.Address {
	address := data.ConvertAddress(*addressDto, order)
	a.aSt.CreateAddress(address)
	return data.ConvertAddressToDto(*address)
}

func (p *productService) CreateProducts(productsDto *[]dto.Product, order model.Order) *[]dto.Product {
	products := data.ConvertProduct(*productsDto, order)
	p.pSt.CreateProducts(products)
	return data.ConvertProductToDto(*products)
}

func (o *orderService) CreateOrder(orderDto *dto.OrderDto) uuid.UUID {
	order := data.ConvertOrder(*orderDto)
	o.oSt.CreateOrder(order)

	//address := data.ConvertAddress(*orderDto, *order)
	o.aS.CreateAddress(&orderDto.Address, *order)

	//products := data.ConvertProduct(*orderDto, *order)
	o.pS.CreateProducts(&orderDto.Products, *order)

	return order.ID
}
