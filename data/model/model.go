package model

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid"` // Foreign key for User
	User        User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
	Roles       []*Role   `json:"roles" gorm:"many2many:user_roles;"`
	Blocked     bool      `json:"is_blocked"`
}

type Role struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name"`
}

type Address struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	OrderID     uuid.UUID `json:"order_id" gorm:"type:uuid"` // Foreign key for Order
	Order       Order     `json:"order" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	HouseNumber string    `json:"house_number"`
	Flat        string    `json:"flat"`
	PostCode    string    `json:"post_code"`
}

type Product struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	ProductID   uuid.UUID `json:"product_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Price       float32   `json:"price"`
	OrderID     uuid.UUID `json:"order_id" gorm:"type:uuid"` // Foreign key for Order
	Order       Order     `json:"order" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
