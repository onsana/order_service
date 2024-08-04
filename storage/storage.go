package storage

import (
	"sync"

	"fmt"

	"github.com/google/uuid"
	"github.com/onsana/order_service/data/model"
	"github.com/onsana/order_service/database"
)

type OrderStorage struct {
	orderM sync.Mutex
}

type ProductStorage struct {
	productM sync.Mutex
}

type AddressStorage struct {
	addressM sync.Mutex
}

func NewOrderStorage() *OrderStorage {
	return &OrderStorage{}
}

func NewProductStorage() *ProductStorage {
	return &ProductStorage{}
}

func NewAddressStorage() *AddressStorage {
	return &AddressStorage{}
}

func (s *OrderStorage) CreateOrder(order *model.Order) model.Order {
	database.DB.Db.Create(order)
	return *order
}

func (s *OrderStorage) GetAllOrders() []model.Order {
	var orders []model.Order
	database.DB.Db.Find(orders)
	return orders
}

func (s *OrderStorage) GetOrderById(id uuid.UUID) (model.Order, error) {
	var order model.Order
	result := database.DB.Db.First(&order, "id = ?", id)
	if result.Error != nil {
		return order, result.Error
	}
	return order, nil
}

func (s *ProductStorage) CreateProducts(products *[]model.Product) ([]model.Product, error) {
	tx := database.DB.Db.Create(products)
	if tx.Error != nil {
		return nil, fmt.Errorf("Error during saving products with ids: %v ", products)
	}
	//TODO remove returned *products
	return *products, nil
}

func (s *AddressStorage) CreateAddress(address *model.Address) model.Address {
	database.DB.Db.Create(address)
	return *address
}
