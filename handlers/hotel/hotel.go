package handlers

import (
	"net/http"
	"subsea/data"

	"github.com/labstack/echo/v4"
)

type Hotels struct {
	db *data.HotelsDB
}

func NewHotels(db *data.HotelsDB) *Hotels {
	return &Hotels{db: db}
}

func (h *Hotels) ListAll(c echo.Context) error {
	c.Echo().Logger.Debug("Hello world")
	return c.String(http.StatusOK, "Hello, World! List All")
}
