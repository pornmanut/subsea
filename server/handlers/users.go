package handlers

import (
	"fmt"
	"net/http"
	"subsea/data"
	"subsea/models"
	"subsea/pwd"
	"subsea/webtoken"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// UserHandler is user handlers
type UserHandler struct {
	v       *data.Validation
	userDB  *data.UserDB
	hotelDB *data.HotelDB
	b       *pwd.Bcrypt
	jwt     *webtoken.JWT
}

// NewUsers is constructor
func NewUsers(db *data.Database, v *data.Validation, b *pwd.Bcrypt, jwt *webtoken.JWT) *UserHandler {
	return &UserHandler{userDB: db.UserDB, hotelDB: db.HotelDB, v: v, b: b, jwt: jwt}
}

// RegisterUser handlers
func (u *UserHandler) RegisterUser(c echo.Context) error {
	c.Echo().Logger.Debug("Register")

	regisUser := c.Get("user").(models.User)

	//check email and username is already exist
	findUser, err := u.userDB.FindOne(
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

	err = u.userDB.Add(regisUser)

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
	user, err := u.userDB.FindOne(bson.M{"username": login.Username})

	if err != nil {
		return c.JSON(http.StatusNotFound, "Not found user")
	}
	// compare password
	if u.b.Compare(login.Password, user.Password) {
		//TODO: jwt generate token

		token, err := u.jwt.CreateToken(user.Username)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, fmt.Sprintf("Token %s", token))
	}
	return c.JSON(http.StatusUnauthorized, "Password doesn't match")
}
