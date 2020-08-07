package main

import (
	"context"
	"net/http"
	"subsea/data"
	handlers "subsea/handlers"
	"subsea/pwd"
	"subsea/webtoken"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind Address for the server")
var dbAddress = env.String("DB_ADDRESS", false, "mongodb://localhost:27017", "Database server Address")
var jwtSecret = env.String("JWT_SECRET", false, "cat", "Secret for jwt")

func main() {
	// parse environment

	err := env.Parse()

	if err != nil {
		panic(err)
	}
	//create new servr
	e := echo.New()

	// create new vlidate
	v := data.NewValidation()
	b := pwd.NewBcrypt(16)
	j := webtoken.NewJWT(6*time.Hour, *jwtSecret)
	middlewareAuth := handlers.NewMiddlewareAuth(j)
	// setting up new log
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := data.ConnectMongoServer(ctx, *dbAddress)
	defer client.Disconnect(ctx)

	db := data.NewDatabase(client, "subsea")

	if err != nil {
		e.Logger.Fatal(err)
	}

	hotelH := handlers.NewHotels(*db.HotelDB)
	userH := handlers.NewUsers(v, db.UserDB, b, j)

	e.Logger.SetLevel(log.DEBUG)

	// basic handler

	e.GET("/hotels", hotelH.ListHotels)
	// e.GET("/hotels/search", nil)
	// e.GET("/hotel/:name", nil)
	// e.POST("/hotel", nil)

	e.POST("/register", userH.RegisterUser, userH.MiddlewareValidateUser)
	e.POST("/login", userH.LoginUser, userH.MiddlewareValidateLogin)
	e.GET("/secret", hello, middlewareAuth)

	// serve server on port
	e.Logger.Fatal(e.Start(*bindAddress))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
