package handlers

import (
	"fmt"
	"net/http"
	"subsea/data"
	"subsea/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// Hotels is hotel handlers
type Hotels struct {
	hotelDB *data.HotelDB
	userDB  *data.UserDB
	v       *data.Validation
}

// NewHotels is constrctor
func NewHotels(db *data.Database, v *data.Validation) *Hotels {
	return &Hotels{hotelDB: db.HotelDB, userDB: db.UserDB, v: v}
}

// TODO:
// middleware extract username
// update username
// update hotel

// Booking booking user for hotel
func (h *Hotels) Booking(c echo.Context) error {
	c.Echo().Logger.Debug("Booking")

	// get request
	tokenDetail := c.Get("myuser").(models.UserTokenDetails)
	hotelName := c.Param("name")
	user, err := h.userDB.FindOne(bson.M{"username": tokenDetail.Username})
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	hotels, err := h.hotelDB.FindOne(bson.M{"name": hotelName})
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	hotels.Booking = hotels.Booking + 1
	user.Bookings = append(user.Bookings, models.Booking(hotels.ID.String()))
	// TODO: must change to objectID
	err = h.hotelDB.ReplaceOne(bson.M{"name": hotelName}, *hotels)
	if err != nil {
		return err
	}
	err = h.userDB.ReplaceOne(bson.M{"username": tokenDetail.Username}, *user)
	if err != nil {
		return err
	}
	fmt.Println(user, hotels)
	return c.JSON(http.StatusOK, user)
}

// ListHotels list all hotel in database
func (h *Hotels) ListHotels(c echo.Context) error {
	c.Echo().Logger.Debug("ListHotels")
	hotels, err := h.hotelDB.Find(bson.M{})
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
	findHotel, err := h.hotelDB.FindOne(bson.M{"name": hotel.Name})

	if err != nil {
		c.Echo().Logger.Error(err)
	}
	// TODO: check no document in result

	if findHotel != nil {
		return c.JSON(http.StatusConflict, `{"message": already exist }`)
	}

	err = h.hotelDB.Add(hotel)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "sucess")
}

// FindOneHotel handlers find one a hotel into database
func (h *Hotels) FindOneHotel(c echo.Context) error {
	c.Echo().Logger.Debug("FindOneHotel")

	name := c.Param("name")
	hotel, err := h.hotelDB.FindOne(bson.M{"name": name})

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
	// for searching
	// we will search only match
	// or use elistic search

	name := c.QueryParam("name")
	detail := c.QueryParam("detail")

	filter := bson.M{}

	if name != "" {
		filter["name"] = name
	}
	if detail != "" {
		filter["detail"] = detail
	}

	result, err := h.hotelDB.Find(filter)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if result == nil {
		return c.JSON(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusOK, result)
}
