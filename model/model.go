package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
}
