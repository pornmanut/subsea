package handlers

import (
	"net/http"
	"subsea/data"

	"github.com/labstack/echo/v4"
)

type Hotels struct {
	db data.HotelsDB
}

func NewHotels(db data.HotelsDB) *Hotels {
	return &Hotels{db: db}
}

func (h *Hotels) ListAll(c echo.Context) error {
	c.Echo().Logger.Debug("Hello world")
	hotels, err := h.db.List()
	if err != nil {
		c.Echo().Logger.Fatal(err)
	}
	return c.JSON(http.StatusOK, hotels)
}
