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

	user := data.User{}

	if err := c.Bind(&user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	c.Echo().Logger.Info(fmt.Sprintf("%+v\n", user))
	return c.NoContent(http.StatusOK)
}
