package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User struct
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email" validate:"required,email"`
	Passwords string    `json:"passwords" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
