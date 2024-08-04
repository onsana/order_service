package storage

import (
	"fmt"
	"github.com/onsana/order_service/data/model"
	"gorm.io/gorm"
	//"sync"
)

type OrderStorage struct {
	//orderM sync.Mutex
	db *gorm.DB
}

type ProductStorage struct {
	//productM sync.Mutex
	db *gorm.DB
}

type AddressStorage struct {
	//addressM sync.Mutex
	db *gorm.DB
}

func NewOrderStorage(db *gorm.DB) *OrderStorage {
	return &OrderStorage{db: db}
}

func NewProductStorage(db *gorm.DB) *ProductStorage {
	return &ProductStorage{db: db}
}

func NewAddressStorage(db *gorm.DB) *AddressStorage {
	return &AddressStorage{db: db}
}

func (s *OrderStorage) CreateOrder(order *model.Order) error {
	tx := s.db.Create(order)
	if tx.Error != nil {
		return fmt.Errorf("Error during saving order: %v ", order)
	}
	return nil
}

func (s *ProductStorage) CreateProducts(products *[]model.Product) error {
	tx := s.db.Create(products)
	if tx.Error != nil {
		return fmt.Errorf("Error during saving products with ids: %v ", products)
	}
	return nil
}

func (s *AddressStorage) CreateAddress(address *model.Address) error {
	tx := s.db.Create(address)
	if tx.Error != nil {
		return fmt.Errorf("Error during saving address: %v ", address)
	}
	return nil
}
