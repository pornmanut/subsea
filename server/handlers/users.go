package handlers

import (
	"fmt"
	"net/http"
	"subsea/data"
	"subsea/models"
	"subsea/pwd"

	"github.com/labstack/echo/v4"
)

// UserHandler is user handlers
type UserHandler struct {
	v  *data.Validation
	db *data.UserDB
	b  *pwd.Bcrypt
}

// NewUsers is constructor
func NewUsers(v *data.Validation, db *data.UserDB, b *pwd.Bcrypt) *UserHandler {
	return &UserHandler{db: db, v: v, b: b}
}

// Register handlers
func (u *UserHandler) Register(c echo.Context) error {
	c.Echo().Logger.Debug("Register")

	user := c.Get("user").(models.User)
	hash, err := u.b.Hash(user.Password)
	if err != nil {
		return err
	}
	// setting new password
	user.Password = hash

	err = u.db.Add(user)

	if err != nil {
		return err
	}
	c.Echo().Logger.Info("Insert User:", fmt.Sprintf("%+v\n", user))
	return c.JSON(http.StatusOK, "success")
}

//TODO: list
//TODO: delete
//TODO: username
