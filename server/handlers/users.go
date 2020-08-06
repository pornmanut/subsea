package handlers

import (
	"fmt"
	"net/http"
	"subsea/data"

	"github.com/labstack/echo/v4"
)

// UserHandler is user handlers
type UserHandler struct {
	db *data.UserDB
}

// NewUsers is constructor
func NewUsers(db *data.UserDB) *UserHandler {
	return &UserHandler{db: db}
}

// Register handlers
func (u *UserHandler) Register(c echo.Context) error {
	c.Echo().Logger.Debug("Register")

	user := data.User{}

	if err := c.Bind(&user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	c.Echo().Logger.Info(fmt.Sprintf("%+v\n", user))
	return c.NoContent(http.StatusOK)
}
