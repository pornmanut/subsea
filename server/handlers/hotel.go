package handlers

import (
	"net/http"
	"subsea/data"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Hotels struct {
	db data.HotelDB
}

func NewHotels(db *data.Database) *Hotels {

	return &Hotels{db: *db.HotelDB}
}

// ListHotels list all hotel in database
func (h *Hotels) ListHotels(c echo.Context) error {
	c.Echo().Logger.Debug("ListHotels")
	hotels, err := h.db.Find(bson.M{})
	if err != nil {
		c.Echo().Logger.Fatal(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, hotels)
}

func (h *Hotels) GetHotel(c echo.Context) error {
	c.Echo().Logger.Debug("GetHotel")

	name := c.Param("name")
	hotels, err := h.db.Find(bson.M{"name": name})

	if err != nil {
		c.Echo().Logger.Fatal(err)
	}
	return c.JSON(http.StatusOK, hotels)
}
