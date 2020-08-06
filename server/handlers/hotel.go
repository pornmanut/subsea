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

// func (h *CustomerHandler) GetCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) SaveCustomer(c echo.Context) error {
// 	customer := Customer{}

// 	if err := c.Bind(&customer); err != nil {
// 		return c.NoContent(http.StatusBadRequest)
// 	}

// 	if err := h.DB.Save(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	if err := c.Bind(&customer); err != nil {
// 		return c.NoContent(http.StatusBadRequest)
// 	}

// 	if err := h.DB.Save(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	if err := h.DB.Delete(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.NoContent(http.StatusNoContent)
// }
