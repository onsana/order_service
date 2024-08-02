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

type OrderService struct {
	oSt orderStorage
	aSt addressStorage
	pSt productStorage
}

func NewService(oSt orderStorage, aSt addressStorage, pSt productStorage) *OrderService {
	return &OrderService{oSt: oSt, aSt: aSt, pSt: pSt}
}

func (o *OrderService) CreateOrder(orderDto *dto.OrderDto) uuid.UUID {
	order := data.ConvertOrder(*orderDto)
	o.oSt.CreateOrder(order)

	address := data.ConvertAddress(*orderDto, *order)
	o.aSt.CreateAddress(address)

	products := data.ConvertProduct(*orderDto, *order)
	o.pSt.CreateProducts(products)

	return order.ID
}
