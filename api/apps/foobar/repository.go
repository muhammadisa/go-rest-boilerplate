package foobar

import (
	"github.com/muhammadisa/go-rest-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

// Repository interface for user
type Repository interface {
	Fetch() ([]models.Foobar, error)
	Create(foobar *models.Foobar) error
	Update(foobar *models.Foobar) error
	Delete(id uuid.UUID) error
}
