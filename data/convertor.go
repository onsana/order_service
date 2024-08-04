package data

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/onsana/order_service/data/dto"
	"github.com/onsana/order_service/data/model"
	"io"
	"log"
	"os"
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

func convertProductMockDtoToProduct(mockDto dto.ProductMockDto) (dto.Product, error) {
	productID, err := uuid.Parse(mockDto.ProductID)
	if err != nil {
		return dto.Product{}, fmt.Errorf("invalid UUID format in covertor: %w", err)
	}

	return dto.Product{
		ProductID:   productID,
		ProductName: mockDto.ProductName,
		Quantity:    0,
		Price:       float32(mockDto.Price),
	}, nil
}

func CreateProductMock() map[uuid.UUID]dto.Product {
	file, err := os.Open("/Users/bigmag/Documents/service-order/order_service/data/dto/productMock.json")
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(file)

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var products []dto.ProductMockDto
	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	productMap := make(map[uuid.UUID]dto.Product)
	for _, product := range products {
		productID, err := uuid.Parse(product.ProductID)
		if err != nil {
			log.Printf("Invalid UUID format: %v", err)
			continue
		}
		productDto, _ := convertProductMockDtoToProduct(product)
		productMap[productID] = productDto
	}
	return productMap
}
