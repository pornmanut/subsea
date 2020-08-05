package main

import (
	"net/http"

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

	//create new servr
	e := echo.New()

	// basic handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// serve server on port
	e.Logger.Fatal(e.Start(*bindAddress))
}
