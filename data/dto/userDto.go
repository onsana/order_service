package dto

import "github.com/google/uuid"

type UserDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"user_name"`
	PhoneNumber string    `json:"phone_number"`
	Roles       []string  `json:"roles"`
	Blocked     bool      `json:"is_blocked"`
}

type UserMockDto struct {
	ID          string   `json:"id"`
	Name        string   `json:"user_name"`
	PhoneNumber string   `json:"phone_number"`
	Roles       []string `json:"roles"`
	Blocked     bool     `json:"is_blocked"`
}
