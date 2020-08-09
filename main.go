package main

import (
	"context"
	"fmt"
	"net/http"
	"subsea/data"
	handlers "subsea/handlers"
	"subsea/pwd"
	"subsea/webtoken"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("PORT", false, "8080", "Bind Address for the server")

// UPDATE GIT
// IN ENV MUST BE THE NAME OF CONTAINER FOR CONNECT such as mogodb://mongoDB:27017
var dbAddress = env.String("DB_ADDRESS", false, "mongodb://localhost:27017", "Database server Address")
var dbName = env.String("DB_NAME", false, "subsea", "Database Name")
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
	fmt.Println("DB address", *dbAddress)
	fmt.Println("DB Name", *dbName)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	client, err := data.ConnectMongoServer(ctx, *dbAddress)
	defer client.Disconnect(ctx)

	if err != nil {
		fmt.Println("Can't connect to database")
		panic(err)
	}

	db := data.NewDatabase(client, *dbName)

	if err != nil {
		e.Logger.Fatal(err)
	}

	hotelH := handlers.NewHotels(db, v)
	userH := handlers.NewUsers(db, v, b, j)

	e.Logger.SetLevel(log.DEBUG)

	// basic handler

	// add cors to public api.
	// allow all
	e.Use(middleware.CORS())

	// POST hotel
	e.POST("/hotels", hotelH.NewHotels, hotelH.MiddlewareValidateHotel)
	e.GET("/hotels/booking/:name", hotelH.Booking, middlewareAuth)

	// GET hotel
	e.GET("/hotels", hotelH.SearchHotel)
	e.GET("/hotels/:name", hotelH.FindOneHotel)
	// e.GET("/hotels/search", hotelH.SearchHotel)

	// POST user
	e.POST("/register", userH.RegisterUser, userH.MiddlewareValidateUser)
	e.POST("/login", userH.LoginUser, userH.MiddlewareValidateLogin)

	e.GET("/users", userH.ListUser, middlewareAuth)
	e.GET("/users/booking", userH.ShowBooking, middlewareAuth)
	e.GET("/secret", hello, middlewareAuth)

	// serve server on port
	e.Logger.Fatal(e.Start(":" + *bindAddress))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
