package handlers

import (
	"net/http"
	"subsea/data"

	"github.com/labstack/echo/v4"
)

type Hotels struct {
	db   data.HotelsDB
	size int
}

func NewHotels(db data.HotelsDB) *Hotels {
	return &Hotels{db: db}
}

// ListHotels list all hotel in database
func (h *Hotels) ListHotels(c echo.Context) error {
	c.Echo().Logger.Debug("ListHotels")
	hotels, err := h.db.List()
	if err != nil {
		c.Echo().Logger.Fatal(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, hotels)
}

func (h *Hotels) GetHotel(c echo.Context) error {
	c.Echo().Logger.Debug("GetHotel")

	name := c.Param("name")
	hotels, err := h.db.Get(name)

	if err != nil {
		c.Echo().Logger.Fatal(err)
	}
	return c.JSON(http.StatusOK, hotels)
}
