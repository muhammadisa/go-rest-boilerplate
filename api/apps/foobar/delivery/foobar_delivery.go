package delivery

import (
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"

	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-rest-boilerplate/api/apps/foobar"
	"github.com/muhammadisa/go-rest-boilerplate/api/models"
	"github.com/muhammadisa/go-rest-boilerplate/api/utils"
)

// FoobarDelivery struct
type FoobarDelivery struct {
	foobarUsecase foobar.Usecase
}

// NewFoobarDelivery function
func NewFoobarDelivery(e *echo.Group, usecase foobar.Usecase) {
	delivery := &FoobarDelivery{
		foobarUsecase: usecase,
	}
	e.GET("/foobars/", delivery.Fetch)
	e.POST("/foobar/", delivery.Create)
	e.PATCH("/foobar/update/:id", delivery.Update)
	e.DELETE("/foobar/delete/:id", delivery.Delete)
}

var model = models.Foobar{}

// Fetch all foobars
func (foobarDeliveries *FoobarDelivery) Fetch(c echo.Context) error {
	var err error

	// err = auth.JWTTokenValidate(c)
	// if err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, utils.Responser{
	// 		StatusCode: http.StatusUnprocessableEntity,
	// 		Message:    err.Error(),
	// 		Data:       nil,
	// 	})
	// }

	foobars, err := foobarDeliveries.foobarUsecase.Fetch()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"status_code": strconv.Itoa(http.StatusUnprocessableEntity),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	return c.JSON(http.StatusOK, utils.Responser{
		StatusCode: http.StatusOK,
		Message:    "Retrieve foobar success",
		Data:       foobars,
	})
}

// Create new foobar
func (foobarDeliveries *FoobarDelivery) Create(c echo.Context) error {
	var err error
	var foobar models.Foobar

	err = c.Bind(&foobar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	err = c.Validate(foobar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	err = foobarDeliveries.foobarUsecase.Create(&foobar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	return c.JSON(http.StatusOK, utils.Responser{
		StatusCode: http.StatusOK,
		Message:    "Create foobar successfully",
		Data:       "Created foobar",
	})
}

// Update existing foobar
func (foobarDeliveries *FoobarDelivery) Update(c echo.Context) error {
	var err error
	var foobar models.Foobar

	err = c.Bind(&foobar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	err = c.Validate(foobar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	err = foobarDeliveries.foobarUsecase.Update(&foobar)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	return c.JSON(http.StatusOK, utils.Responser{
		StatusCode: http.StatusOK,
		Message:    "Update foobar successfully",
		Data:       "Updated foobar",
	})
}

// Delete delete account permanently
func (foobarDeliveries *FoobarDelivery) Delete(c echo.Context) error {
	var err error

	foobarID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	err = foobarDeliveries.foobarUsecase.Delete(foobarID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status_code": strconv.Itoa(http.StatusBadRequest),
			"message":     err.Error(),
			"data":        "nil",
		})
	}
	return c.JSON(http.StatusOK, utils.Responser{
		StatusCode: http.StatusOK,
		Message:    "Foobar data deleted",
		Data:       "OK",
	})
}
