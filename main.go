package main

import (
	"subsea/data"
	handlers "subsea/handlers/hotel"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind Address for the server")

func main() {
	// parse environment

	err := env.Parse()

	if err != nil {
		panic(err)
	}

	// setting up new log

	db := data.NewHotelDB()
	hh := handlers.NewHotels(db)

	//create new servr
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// basic handler
	e.GET("/", hh.ListAll)
	// serve server on port
	e.Logger.Fatal(e.Start(*bindAddress))
}

// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }
