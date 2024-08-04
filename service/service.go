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
	CreateOrder(order *model.Order) model.Order
	GetAllOrders() []model.Order
}

type addressStorage interface {
	CreateAddress(address *model.Address) model.Address
}

type productStorage interface {
	CreateProducts(order *[]model.Product) ([]model.Product, error)
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

func (a *addressService) CreateAddress(addressDto *dto.Address, order model.Order) *dto.Address {
	address := data.ConvertAddress(*addressDto, order)
	a.aSt.CreateAddress(address)
	return data.ConvertAddressToDto(*address)
}

func (p *productService) CreateProducts(productsDto *[]dto.Product, order model.Order) (*[]dto.Product, error) {
	products := data.ConvertProduct(*productsDto, order)
	_, err := p.pSt.CreateProducts(products)
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
	o.oSt.CreateOrder(order)

	o.aS.CreateAddress(&orderDto.Address, *order)

	_, err = o.pS.CreateProducts(validatedProducts, *order)
	if err != nil {
		return uuid.Nil, err
	}

	return order.ID, nil
}
func (o *orderService) GetAllOrders() []model.Order {
	return o.oSt.GetAllOrders()
}
