package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Foobar struct
type Foobar struct {
	ID            uuid.UUID `db:"id" json:"id"`
	FoobarContent string    `json:"foobar_content" validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
