package delivery

import (
	"net/http"
	"strconv"

	"github.com/muhammadisa/go-rest-boilerplate/api/utils"
	uuid "github.com/satori/go.uuid"

	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-rest-boilerplate/api/apps/user"
	"github.com/muhammadisa/go-rest-boilerplate/api/models"
)

// UserDelivery struct
type UserDelivery struct {
	userUsecase user.Usecase
}

// NewUserDelivery function
func NewUserDelivery(e *echo.Group, usecase user.Usecase) {
	delivery := &UserDelivery{
		userUsecase: usecase,
	}
	e.POST("/user/login/", delivery.Login)
	e.POST("/user/register/", delivery.Register)
	e.DELETE("/user/close/account/:id", delivery.CloseAccount)
}

var model = models.User{}

// Login and auth user
func (userDeliveries UserDelivery) Login(c echo.Context) error {
	var err error
	var user models.User

	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	err = c.Validate(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	_, authentic, err := userDeliveries.userUsecase.Login(&user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"status_code": strconv.Itoa(http.StatusUnauthorized),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	return c.JSON(http.StatusOK, utils.Responser{
		StatusCode: http.StatusOK,
		Message:    "Login Successfully",
		Data:       authentic,
	})
}

// Register new user
func (userDeliveries *UserDelivery) Register(c echo.Context) error {
	var err error
	var user models.User

	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	err = c.Validate(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	err = userDeliveries.userUsecase.Register(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	return c.JSON(http.StatusOK, utils.Responser{
		StatusCode: http.StatusOK,
		Message:    "Registered Successfully",
		Data:       "Successfully registered",
	})
}

// CloseAccount delete account permanently
func (userDeliveries *UserDelivery) CloseAccount(c echo.Context) error {
	var err error

	userID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	err = userDeliveries.userUsecase.Delete(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	return c.JSON(http.StatusOK, utils.Responser{
		StatusCode: http.StatusOK,
		Message:    "User data deleted",
		Data:       "OK",
	})
}
