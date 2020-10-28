package user

import (
	"github.com/muhammadisa/go-rest-boilerplate/api/auth"
	"github.com/muhammadisa/go-rest-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

// Usecase interaface for user
type Usecase interface {
	Login(user *models.User) (*models.User, *auth.Authenticated, error)
	Register(user *models.User) error
	Update(user *models.User) error
	Delete(id uuid.UUID) error
}
