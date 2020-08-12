package main

import (
	"context"
	"net/http"
	"os"
	"subsea/data"
	handlers "subsea/handlers"
	mymiddleware "subsea/middleware"
	"subsea/mongo"
	"subsea/pwd"
	"subsea/webtoken"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("PORT", false, "8080", "Bind Address for the server")
var dbAddress = env.String("DB_ADDRESS", false, "mongodb://localhost:27017", "Database server Address")
var dbName = env.String("DB_NAME", false, "subsea", "Database Name")
var jwtSecret = env.String("JWT_SECRET", false, "cat", "Secret for jwt")

func main() {
	// parse environment
	err := env.Parse()

	if err != nil {
		panic(err)
	}

	// setting up log
	l := hclog.New(hclog.DefaultOptions)

	// get envrioment
	l.Info("Port address:", *bindAddress)
	l.Info("Database address:", *dbAddress)
	l.Info("Database name:", *dbName)
	l.Info("JWT secret:", *jwtSecret)

	// initialize basic setup
	v := data.NewValidation()
	b := pwd.NewBcrypt(16)
	j := webtoken.NewJWT(6*time.Hour, *jwtSecret)

	// initalize mongo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := mongo.ConnectMongoServer(ctx, *dbAddress)
	defer client.Disconnect(ctx)

	if err != nil {
		l.Error("Cannot connect to database", "error", err)
		os.Exit(1)
	}

	// initalize mongo database
	database := client.Database(*dbName)
	hotelDB := mongo.NewHotelDB(database)
	userDB := mongo.NewUserDB(database)

	// get database target to database
	db := data.NewDatabase(userDB, hotelDB)

	if err != nil {
		l.Error("Database", "error", err)
	}

	// initailze handler
	hotelH := handlers.NewHotelsHandler(db, l)
	userH := handlers.NewUsersHandler(db, l, j, b)
	middlewareH := mymiddleware.NewMiddleware(v, j)

	// create new server
	// using echo framework
	e := echo.New()

	// CORS using options for public api
	e.Use(middleware.CORS())

	// route

	//
	// e.GET("/hotels/booking/:name", hotelH.Booking, middlewareAuth)

	// GET hotel
	e.GET("/hotels", hotelH.ListHotels)
	e.POST("/hotels", hotelH.NewHotels, middlewareH.MiddlewareAuth)
	e.GET("/hotels/:name", hotelH.FindOneHotel)
	// e.GET("/hotels/search", hotelH.SearchHotel)

	// POST user
	// e.POST("/register", userH.RegisterUser, userH.MiddlewareValidateUser)
	e.POST("/login", userH.LoginUser, middlewareH.MiddlewareValidateLogin)

	// e.GET("/users", userH.ListUser, middlewareAuth)
	// e.GET("/users/booking", userH.ShowBooking, middlewareAuth)
	e.GET("/secret", hello, middlewareH.MiddlewareAuth)

	// serve server on port
	e.Logger.Fatal(e.Start(":" + *bindAddress))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
