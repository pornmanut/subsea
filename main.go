package main

import (
	"context"
	"subsea/data"
	handlers "subsea/handlers"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind Address for the server")
var dbAddress = env.String("DB_ADDRESS", false, "mongodb://localhost:27017", "Database server Address")

func main() {
	// parse environment

	err := env.Parse()

	if err != nil {
		panic(err)
	}
	//create new servr
	e := echo.New()
	// setting up new log
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := data.ConnectMongoServer(ctx, *dbAddress)
	defer client.Disconnect(ctx)

	db, err := data.NewHotelMongo(client)

	if err != nil {
		e.Logger.Fatal(err)
	}

	hotelH := handlers.NewHotels(db)

	e.Logger.SetLevel(log.DEBUG)

	// basic handler

	e.GET("/hotels", hotelH.ListHotels)
	// e.GET("/hotels/search", nil)
	// e.GET("/hotel/:name", nil)
	// e.POST("/hotel", nil)

	e.POST("/register", nil)

	// serve server on port
	e.Logger.Fatal(e.Start(*bindAddress))
}

// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }
