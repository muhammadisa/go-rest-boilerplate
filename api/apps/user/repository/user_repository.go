package repository

import (
	"errors"
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

	rowsAffected, err := userRepositories.Sess.Select("*").
		From("users").
		Where("email = ?", user.Email).
		Load(&selectedUser)
	if rowsAffected == 0 {
		return nil, nil, errors.New("User not found")
	}
	if err != nil {
		return nil, nil, err
	}
	return selectedUser, &auth.Authenticated{}, nil
}

func (userRepositories *userRepository) Register(user *models.User) error {
	var err error
	// kontol := models.User{
	// 	ID:        uuid.NewV4(),
	// 	Passwords: "$2a$10$9e5Qfi1oz7H78r4QSqR5Cu6W2Oh8.2tXdXHQc10PoAD6OM4d0Ku5a",
	// 	Email:     "asdasdasd@gmail.com",
	// 	CreatedAt: time.Now(),
	// }
	_, err = userRepositories.Sess.InsertInto("users").
		Columns("id", "email", "passwords", "created_at").
		Record(user).
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
			"password":   user.Passwords,
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
