package dto

import (
	"time"

	"github.com/google/uuid"
)

// TODO remove dto from name
type OrderDto struct {
	ID          uuid.UUID `json:"id"`
	UserDto     UserDto   `json:"user,omitempty"`
	Products    []Product `json:"products,omitempty"`
	Address     Address   `json:"address,omitempty"`
	TotalPrice  float32   `json:"totalPrice"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Status      Kind      `json:"status"`
}

type Address struct {
	ID          uuid.UUID `json:"id"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	HouseNumber string    `json:"house_number"`
	Flat        string    `json:"flat"`
	PostCode    string    `json:"post_code"`
}

type Product struct {
	ProductID   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Price       float32   `json:"price"`
}

type ProductMockDto struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
}

type Kind string

const (
	// Pending indicates that the order has been created buyer
	Pending Kind = "pending"

	// Paid indicates that the order has been paid for.
	Paid Kind = "paid"

	// Canceled indicates that the administrator has canceled the order from Pending or Paid status.
	Canceled Kind = "canceled"
)
