package handlers

import (
	"fmt"
	"net/http"
	"subsea/data"
	"subsea/errors"
	"subsea/models"

	hclog "github.com/hashicorp/go-hclog"

	"github.com/labstack/echo/v4"
)

// HotelsHandler is hotel handlers
type HotelsHandler struct {
	db *data.Database
	l  hclog.Logger
}

// NewHotelsHandler is constrctor
func NewHotelsHandler(db *data.Database, l hclog.Logger) *HotelsHandler {
	return &HotelsHandler{db: db, l: l}
}

// TODO:
// middleware extract username
// update username
// update hotel

// Booking booking user for hotel
// func (h *Hotels) Booking(c echo.Context) error {
// 	c.Echo().Logger.Debug("Booking")

// 	// get request
// 	tokenDetail := c.Get("myuser").(models.UserTokenDetails)
// 	hotelName := c.Param("name")
// 	user, err := h.userDB.FindOne(bson.M{"username": tokenDetail.Username})
// 	if err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}
// 	hotels, err := h.hotelDB.FindOne(bson.M{"name": hotelName})
// 	if err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}
// 	if hotels.Booking >= hotels.Max {
// 		return c.JSON(400, "Can't booking no avaliable")
// 	}

// 	hotels.Booking = hotels.Booking + 1
// 	user.Bookings = append(user.Bookings, models.Booking(hotels.ID.Hex()))
// 	// TODO: must change to objectID
// 	err = h.hotelDB.ReplaceOne(bson.M{"name": hotelName}, *hotels)
// 	if err != nil {
// 		return err
// 	}
// 	err = h.userDB.ReplaceOne(bson.M{"username": tokenDetail.Username}, *user)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(user, hotels)
// 	return c.JSON(http.StatusOK, user)
// }

// ListHotels list all hotel in database
func (h *HotelsHandler) ListHotels(c echo.Context) error {
	h.l.Info("HotelHandler", "ListHotels Request")
	hotels, err := h.db.HotelDB.ListAllHotels()

	if hotels == nil {
		h.l.Info("Not found any hotel")
		return c.JSON(
			http.StatusNotFound,
			models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: errors.ErrNoDocuments.Error(),
			},
		)

	}
	if err := handleUnknowError(err); err != nil {
		return c.JSON(err.Code, err)
	}
	h.l.Info("Succesfully", "ListHotel")
	return c.JSON(http.StatusOK, hotels)
}

// NewHotels handlers add a hotel into database
func (h *HotelsHandler) NewHotels(c echo.Context) error {
	hotel := c.Get("hotel").(models.Hotel)

	if err := h.isNameExists(hotel.Name); err != nil {
		return c.JSON(err.Code, err)
	}

	id, err := h.db.HotelDB.CreateHotel(hotel)

	if err := handleUnknowError(err); err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("success with %s", id))
}

// FindOneHotel handlers find one a hotel into database
func (h *HotelsHandler) FindOneHotel(c echo.Context) error {

	name := c.Param("name")
	hotel, err := h.db.HotelDB.FindHotelByName(name)

	if hotel == nil {
		h.l.Info("Not found hotel")
		return c.JSON(
			http.StatusNotFound,
			models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: errors.ErrNoDocuments.Error(),
			},
		)

	}
	if err := handleUnknowError(err); err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, hotel)
}

func (h *HotelsHandler) isNameExists(name string) *models.ErrorResponse {
	h.l.Info("Checking Name", "name", name)
	user := new(models.Hotel)
	user, err := h.db.HotelDB.FindHotelByName(name)
	if err := handleUnknowError(err); err != nil {
		return err
	}

	// user exists
	if user != nil {
		h.l.Error("Conflict", "user", user)
		return &models.ErrorResponse{
			Code:    http.StatusConflict,
			Message: errors.ErrEmailAlreadyExists.Error(),
		}
	}
	return nil
}

// // SearchHotel querry params
// func (h *Hotels) SearchHotel(c echo.Context) error {
// 	c.Echo().Logger.Debug("SearchHotel")

// 	// TODO:
// 	// REFACTOR !!!!!!!!!!!!!!!!!!!!!!
// 	// YOU CODE IS SLOW

// 	name := c.QueryParam("name")
// 	detail := c.QueryParam("detail")
// 	lt := c.QueryParam("lt")
// 	gt := c.QueryParam("gt")

// 	requestTag := []bson.M{}
// 	numOfOptions := 0
// 	if name != "" {
// 		requestTag = append(
// 			requestTag,
// 			bson.M{"name": bson.D{
// 				{"$regex", primitive.Regex{Pattern: name, Options: "i"}},
// 			}})
// 		numOfOptions++
// 	}
// 	if detail != "" {
// 		requestTag = append(
// 			requestTag,
// 			bson.M{"detail": bson.D{
// 				{"$regex", primitive.Regex{Pattern: detail, Options: "i"}},
// 			}})
// 		numOfOptions++
// 	}
// 	if lt != "" {
// 		max, err := strconv.Atoi(lt)
// 		if err != nil {
// 			return err
// 		}
// 		requestTag = append(
// 			requestTag,
// 			bson.M{"price": bson.M{
// 				"$lt": max,
// 			}})
// 		numOfOptions++
// 	}
// 	if gt != "" {
// 		min, err := strconv.Atoi(lt)
// 		if err != nil {
// 			return err
// 		}
// 		requestTag = append(
// 			requestTag,
// 			bson.M{"price": bson.M{
// 				"$gt": min,
// 			}})
// 		numOfOptions++
// 	}
// 	filter := bson.M{"$and": requestTag}

// 	if numOfOptions == 0 {
// 		filter = bson.M{}
// 	}
// 	fmt.Println(filter)
// 	result, err := h.hotelDB.Find(filter)
// 	fmt.Println(result)
// 	if err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}
// 	return c.JSON(http.StatusOK, result)
// }
