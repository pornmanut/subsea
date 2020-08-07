package handlers

import (
	"net/http"
	"subsea/data"
	"subsea/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Hotels struct {
	db *data.HotelDB
	v  *data.Validation
}

func NewHotels(db *data.Database, v *data.Validation) *Hotels {
	return &Hotels{db: db.HotelDB, v: v}
}

// ListHotels list all hotel in database
func (h *Hotels) ListHotels(c echo.Context) error {
	c.Echo().Logger.Debug("ListHotels")
	hotels, err := h.db.Find(bson.M{})
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if hotels == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, hotels)
}

// NewHotels handlers add a hotel into database
func (h *Hotels) NewHotels(c echo.Context) error {
	c.Echo().Logger.Debug("NewHotels")

	hotel := c.Get("hotel").(models.Hotel)
	findHotel, err := h.db.FindOne(bson.M{"name": hotel.Name})

	if err != nil {
		c.Echo().Logger.Error(err)
	}
	// TODO: check no document in result

	if findHotel != nil {
		return c.JSON(http.StatusConflict, `{"message": already exist }`)
	}

	err = h.db.Add(hotel)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "sucess")
}

// FindOneHotel handlers find one a hotel into database
func (h *Hotels) FindOneHotel(c echo.Context) error {
	c.Echo().Logger.Debug("FindOneHotel")

	name := c.Param("name")
	hotel, err := h.db.FindOne(bson.M{"name": name})

	if err != nil {
		c.Echo().Logger.Error(err)
	}

	if hotel == nil {
		return c.JSON(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusOK, hotel)
}

// SearchHotel querry params
func (h *Hotels) SearchHotel(c echo.Context) error {
	c.Echo().Logger.Debug("SearchHotel")

	// TODO:
	// index (I don't know how to use index in mongoDB)
	// for search

	name := c.QueryParam("name")
	detail := c.QueryParam("detail")

	filter := bson.M{}

	if name != "" {
		filter["name"] = name
	}
	if detail != "" {
		filter["detail"] = detail
	}

	result, err := h.db.Find(filter)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if result == nil {
		return c.JSON(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusOK, result)
}
