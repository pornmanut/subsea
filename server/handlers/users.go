package handlers

import (
	"fmt"
	"net/http"
	"subsea/data"

	"github.com/labstack/echo/v4"
)

// UserHandler is user handlers
type UserHandler struct {
	v  *data.Validation
	db *data.UserDB
}

// NewUsers is constructor
func NewUsers(v *data.Validation, db *data.UserDB) *UserHandler {
	return &UserHandler{db: db, v: v}
}

// Register handlers
func (u *UserHandler) Register(c echo.Context) error {
	c.Echo().Logger.Debug("Register")

	user := c.Get("user")

	c.Echo().Logger.Info(fmt.Sprintf("%+v\n", user))
	return c.JSON(http.StatusOK, user)
}
