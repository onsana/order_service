package data

import (
	"fmt"
	"github.com/onsana/order_service/data/dto"
	"github.com/onsana/order_service/data/model"
)

func convertUser(userDto dto.UserDto) model.User {
	user := model.User{
		ID: userDto.ID,
	}
	return user
}

func ConvertAddress(addressDto dto.Address, order model.Order) *model.Address {
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

func ConvertAddressToDto(address model.Address) *dto.Address {
	fmt.Printf("addressdto  2  %s\n", address)
	addressDto := dto.Address{
		ID:          address.ID,
		City:        address.City,
		Street:      address.Street,
		HouseNumber: address.HouseNumber,
		Flat:        address.Flat,
		PostCode:    address.PostCode,
	}
	fmt.Printf("address  2  %s\n", address)
	return &addressDto
}

func ConvertOrder(orderDto dto.OrderDto) *model.Order {
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

func ConvertProduct(productsDto []dto.Product, order model.Order) *[]model.Product {
	var products []model.Product
	for _, p := range productsDto {
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

func ConvertProductToDto(products []model.Product) *[]dto.Product {
	var productsDTO []dto.Product
	for _, p := range products {
		product := dto.Product{
			ProductID:   p.ProductID,
			ProductName: p.ProductName,
			Quantity:    p.Quantity,
			Price:       p.Price,
		}
		productsDTO = append(productsDTO, product)
	}
	return &productsDTO
}
