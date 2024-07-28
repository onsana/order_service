package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID `json:"id"`
	UserID      int       `json:"user_id"`
	Products    []Product `json:"products" gorm:"many2many:order_products;"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	HouseNumber string    `json:"house_number"`
	Flat        string    `json:"flat"`
	PostCode    string    `json:"post_code"`
	TotalPrice  float32   `json:"totalPrice"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Status      Kind      `json:"status"`
}

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"user_name"`
	PhoneNumber string    `json:"phone_number"`
	Roles       []string  `json:"roles"`
	Blocked     bool      `json:"is_blocked"`
}

type Address struct {
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	Flat        string `json:"flat"`
	PostCode    string `json:"post_code"`
}

type Product struct {
	ProductID   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Price       float32   `json:"price"`
}

type Kind string

const (
	// Pending indicates that the order has been created buyer
	Pending Kind = "pending"

	// Paid indicates that the order has been paid for.
	Paid Kind = "paid"

	// Delivered indicates that the order has been delivered to the buyer.
	Delivered Kind = "delivered"

	// Canceled indicates that the administrator has canceled the order from Pending or Paid status.
	Canceled Kind = "canceled"
)
