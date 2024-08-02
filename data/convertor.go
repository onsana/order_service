package data

import (
	"fmt"
	dto2 "github.com/onsana/order_service/data/dto"
	"github.com/onsana/order_service/data/model"
)

func convertUser(userDto dto2.UserDto) model.User {
	user := model.User{
		ID: userDto.ID,
	}
	return user
}

func ConvertAddress(orderDto dto2.OrderDto, order model.Order) *model.Address {
	addressDto := orderDto.Address
	fmt.Printf("addressdto  2  %s\n", addressDto)
	address := model.Address{
		ID:          addressDto.ID,
		OrderID:     order.ID,
		Order:       order,
		City:        addressDto.City,
		Street:      addressDto.Street,
		HouseNumber: addressDto.HouseNumber,
		Flat:        addressDto.Flat,
		PostCode:    addressDto.PostCode,
	}
	fmt.Printf("address  2  %s\n", address)
	return &address
}

func ConvertOrder(orderDto dto2.OrderDto) *model.Order {
	user := convertUser(orderDto.UserDto)
	order := &model.Order{
		ID:          orderDto.ID,
		UserID:      user.ID,
		User:        user,
		TotalPrice:  orderDto.TotalPrice,
		CreatedAt:   orderDto.CreatedAt,
		UpdatedAt:   orderDto.UpdatedAt,
		Description: orderDto.Description,
		Status:      model.Kind(orderDto.Status),
	}
	return order
}

func ConvertProduct(orderDto dto2.OrderDto, order model.Order) *[]model.Product {
	var products []model.Product
	for _, p := range orderDto.Products {
		product := model.Product{
			ProductID:   p.ProductID,
			ProductName: p.ProductName,
			Quantity:    p.Quantity,
			Price:       p.Price,
			OrderID:     order.ID,
		}
		products = append(products, product)
	}

	return &products
}
