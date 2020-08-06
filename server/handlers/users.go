package handlers

import (
	"fmt"
	"net/http"
	"subsea/data"
	"subsea/models"
	"subsea/pwd"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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

// RegisterUser handlers
func (u *UserHandler) RegisterUser(c echo.Context) error {
	c.Echo().Logger.Debug("Register")

	regisUser := c.Get("user").(models.User)

	//check email and username is already exist
	findUser, err := u.db.FindOne(
		bson.M{"$or": []interface{}{
			bson.M{"username": regisUser.Username},
			bson.M{"email": regisUser.Email},
		}})

	if findUser != nil {
		return c.JSON(http.StatusConflict, `{"message": already exist }`)
	}

	hash, err := u.b.Hash(regisUser.Password)
	if err != nil {
		return err
	}
	// setting new password
	regisUser.Password = hash

	err = u.db.Add(regisUser)

	if err != nil {
		return err
	}
	c.Echo().Logger.Info("Insert User:", fmt.Sprintf("%+v\n", regisUser))
	return c.JSON(http.StatusOK, "success")
}

//TODO: list
//TODO: delete
//TODO: username

// LoginUser handlers is a login handlers to get models.login
// and compare password and hash return jwt token for store in front-side
func (u *UserHandler) LoginUser(c echo.Context) error {
	c.Echo().Logger.Debug("Login")

	login := c.Get("login").(models.Login)

	// find user for login
	user, err := u.db.FindOne(bson.M{"username": login.Username})

	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	// compare password
	if u.b.Compare(login.Password, user.Password) {
		//TODO: jwt generate token
		return c.JSON(http.StatusOK, "match")
	}
	return c.NoContent(http.StatusUnauthorized)
}
