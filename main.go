package main

import (
	"subsea/data"
	handlers "subsea/handlers/hotel"

	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
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
	l := hclog.Default()

	db := data.NewHotelDB(l)
	hh := handlers.NewHotels(l, db)

	//create new servr
	e := echo.New()

	// basic handler
	e.GET("/", hh.ListAll)
	// serve server on port
	e.Logger.Fatal(e.Start(*bindAddress))
}

// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }
