package api

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-rest-boilerplate/api/routes"

	"github.com/gocraft/dbr/dialect"
	"github.com/gocraft/dbr/v2"
	"github.com/joho/godotenv"
	"github.com/muhammadisa/go-rest-boilerplate/api/utils/errhandler"
	"github.com/muhammadisa/godbconn"
)

// Run start server & connecting to db
func Run() {

	// Loading .env file
	err := godotenv.Load()
	errhandler.HandleError(err, true)

	// Load database credential env and use it
	db, err := godbconn.DBCred{
		DBDriver:   os.Getenv("DB_DRIVER"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}.Connect()
	errhandler.HandleError(err, true)
	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.MySQL,
	}
	conn.SetMaxOpenConns(10)
	session := conn.NewSession(nil)
	session.Begin()

	// Starting echo web framework
	routes.RouteConfigs{
		EchoData:  echo.New(),
		Sess:      session,
		APISecret: os.Getenv("API_SECRET"),
		Version:   "v1",
		Port:      os.Getenv("HTTP_PORT"),
		Origins:   strings.Split(os.Getenv("ORIGINS"), ","),
	}.NewHTTPRoute()

}
