package handlers

import (
	"net/http"
	"subsea/data"

	"github.com/labstack/echo/v4"
)

// UserHandler is user handlers
type UserHandler struct {
	db data.UserDB
}

// NewUsers is constructor
func NewUsers(db data.UserDB) *UserHandler {
	return &UserHandler{db: db}
}

func (u *UserHandler) Register(c echo.Context) error {
	c.Echo().Logger.Debug("Register")

	return c.NoContent(http.StatusOK)
}
