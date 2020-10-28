package usecase

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/muhammadisa/go-rest-boilerplate/api/apps/user"
	"github.com/muhammadisa/go-rest-boilerplate/api/auth"
	"github.com/muhammadisa/go-rest-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

type userUsecase struct {
	userRepository user.Repository
}

// NewUserUsecase function
func NewUserUsecase(userRepository user.Repository) user.Usecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (userUsecases userUsecase) Login(
	user *models.User,
) (
	*models.User,
	*auth.Authenticated,
	error,
) {
	selectedUser, authenticated, err := userUsecases.userRepository.Login(user)
	err = auth.VerifyPassword(selectedUser.Password, user.Password)
	if err != nil {
		return nil, nil, errors.New("Email or Password is incorrect")
	}
	token, refresh, err := auth.GenerateToken(selectedUser.ID, jwt.MapClaims{})
	if err != nil {
		return nil, nil, err
	}
	authenticated.AccessToken = token
	authenticated.RefreshToken = refresh
	return user, authenticated, nil
}

func (userUsecases userUsecase) Register(user *models.User) error {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.ID = uuid.NewV4()
	user.Password = string(hashedPassword)
	err = userUsecases.userRepository.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func (userUsecases userUsecase) Update(user *models.User) error {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.UpdatedAt = time.Now()
	user.Password = string(hashedPassword)
	err = userUsecases.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (userUsecases userUsecase) Delete(id uuid.UUID) error {
	return userUsecases.userRepository.Delete(id)
}
