package repository

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/gocraft/dbr/v2"
	"github.com/muhammadisa/go-rest-boilerplate/api/apps/user"
	"github.com/muhammadisa/go-rest-boilerplate/api/auth"
	"github.com/muhammadisa/go-rest-boilerplate/api/models"
)

type userRepository struct {
	Sess *dbr.Session
}

// NewUserRepository function
func NewUserRepository(sess *dbr.Session) user.Repository {
	return &userRepository{
		Sess: sess,
	}
}

func (userRepositories *userRepository) Login(user *models.User) (*models.User, *auth.Authenticated, error) {
	var err error
	var selectedUser *models.User

	_, err = userRepositories.Sess.Select("*").
		From("users").
		Where("email = ?", user.Email).
		Load(&selectedUser)
	if err != nil {
		return nil, nil, err
	}
	return selectedUser, &auth.Authenticated{}, nil
}

func (userRepositories *userRepository) Register(user *models.User) error {
	var err error

	_, err = userRepositories.Sess.InsertInto("users").
		Columns("id", "email", "password", "created_at").
		Record(&user).
		Exec()
	if err != nil {
		return err
	}
	return nil
}

func (userRepositories *userRepository) Update(user *models.User) error {
	var err error

	_, err = userRepositories.Sess.Update("users").
		Where("id = ?", user.ID.String()).
		SetMap(map[string]interface{}{
			"email":      user.Email,
			"password":   user.Password,
			"updated_at": time.Now(),
		}).
		Exec()
	if err != nil {
		return err
	}
	return nil
}

func (userRepositories *userRepository) Delete(id uuid.UUID) error {
	var err error

	_, err = userRepositories.Sess.DeleteFrom("users").
		Where("id = ?", id.String()).
		Exec()
	if err != nil {
		return err
	}
	return nil
}
