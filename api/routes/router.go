package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/muhammadisa/go-rest-boilerplate/api/middlewares"

	_userApi "github.com/muhammadisa/go-rest-boilerplate/api/apps/user/delivery"
	_userRepo "github.com/muhammadisa/go-rest-boilerplate/api/apps/user/repository"
	_userUsecase "github.com/muhammadisa/go-rest-boilerplate/api/apps/user/usecase"
	"github.com/muhammadisa/goredisku"

	_foobarApi "github.com/muhammadisa/go-rest-boilerplate/api/apps/foobar/delivery"
	_foobarRepo "github.com/muhammadisa/go-rest-boilerplate/api/apps/foobar/repository"
	_foobarUsecase "github.com/muhammadisa/go-rest-boilerplate/api/apps/foobar/usecase"

	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	customValidator "github.com/muhammadisa/go-rest-boilerplate/api/validator"
	"gopkg.in/go-playground/validator.v9"
)

// Routes struct
type Routes struct {
	Echo         *echo.Echo
	RestGroup    *echo.Group
	AuthGroup    *echo.Group
	CacheCommand *goredisku.GoRedisKu
	Sess         *dbr.Session
}

// RouteConfigs struct
type RouteConfigs struct {
	EchoData     *echo.Echo
	Sess         *dbr.Session
	CacheCommand *goredisku.GoRedisKu
	APISecret    string
	Version      string
	Port         string
	Origins      []string
}

// NewHTTPRoute echo route initialization
func (rc RouteConfigs) NewHTTPRoute() {
	// Initialize route configs
	restful := rc.EchoData.Group(fmt.Sprintf("api/%s", rc.Version))
	auth := rc.EchoData.Group(fmt.Sprintf("api/auth/%s", rc.Version))
	handler := &Routes{
		Echo:         rc.EchoData,
		RestGroup:    restful,
		AuthGroup:    auth,
		CacheCommand: rc.CacheCommand,
		Sess:         rc.Sess,
	}
	handler.Echo.Validator = &customValidator.
		CustomValidator{Validator: validator.New()}
	handler.setupMiddleware(rc.APISecret, rc.Origins)
	handler.setInitRoutes()

	// Internal routers
	handler.initUserRoutes()
	handler.initFoobarRoutes()

	// Starting Echo Server
	log.Fatal(handler.Echo.Start(rc.Port))
}

func (r *Routes) setupMiddleware(apiSecret string, origins []string) {
	// main middleware
	r.RestGroup.Use(middlewares.InitMiddleware().JWT)
	r.Echo.Use(middlewares.InitMiddleware().CORS)
	r.Echo.Use(middleware.Recover())
	r.Echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	r.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodOptions,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
	}))
}

func (r *Routes) setInitRoutes() {
	r.Echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status_code": strconv.Itoa(http.StatusOK),
			"message":     "Server is started",
		})
	})
}

// Create route initialization function here
func (r *Routes) initUserRoutes() {
	userRepo := _userRepo.NewUserRepository(r.Sess)
	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	_userApi.NewUserDelivery(r.AuthGroup, userUsecase)
}

// Create route initialization function here
func (r *Routes) initFoobarRoutes() {
	foobarRepo := _foobarRepo.NewFoobarRepository(r.Sess)
	foobarUsecase := _foobarUsecase.NewFoobarUsecase(foobarRepo, r.CacheCommand)
	_foobarApi.NewFoobarDelivery(r.RestGroup, foobarUsecase)
}
