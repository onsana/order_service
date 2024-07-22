package model

import (
	"go/constant"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      int           `json:"user_id"`
	Products    []Product     `json:"products"`
	Address     Address       `gorm:"references:Address" json:"address"`
	TotalPrice  float32       `json:"totalPrice"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Description string        `json:"description"`
	Status      constant.Kind `json:"status"`
	User        User          `gorm:"foreignKey:UserID"`
}

type User struct {
	gorm.Model
	Name        string   `json:"user_name"`
	PhoneNumber string   `json:"phone_number"`
	Roles       []string `json:"roles"`
	Blocked     bool     `json:"is_blocked"`
}

type Address struct {
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	Flat        string `json:"flat"`
	PostCode    string `json:"post_code"`
}

type Product struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
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
