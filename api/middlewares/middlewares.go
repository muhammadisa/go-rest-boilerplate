package middlewares

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-rest-boilerplate/api/auth"
)

// Middlewares struct
type Middlewares struct{}

// CORS avoid CORS Error
func (m *Middlewares) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// JWT jwt auth middleware
func (m *Middlewares) JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := auth.JWTTokenValidate(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"status_code": strconv.Itoa(http.StatusUnauthorized),
				"message":     "Unauthorized",
			})
			return nil
		}
		return next(c)
	}
}

// InitMiddleware initialize middleware
func InitMiddleware() *Middlewares {
	return &Middlewares{}
}
