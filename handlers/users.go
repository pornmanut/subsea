package handlers

import (
	"fmt"
	"net/http"
	"subsea/data"
	"subsea/errors"
	"subsea/models"
	"subsea/pwd"
	"subsea/webtoken"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
)

// UserHandler is user handlers
type UserHandler struct {
	db  *data.Database
	l   hclog.Logger
	b   *pwd.Bcrypt
	jwt *webtoken.JWT
}

// NewUsersHandler is constructor
func NewUsersHandler(db *data.Database, l hclog.Logger, jwt *webtoken.JWT, b *pwd.Bcrypt) *UserHandler {
	return &UserHandler{db: db, l: l, b: b, jwt: jwt}
}

// //ShowBooking handler to show booking
// func (u *UserHandler) ShowBooking(c echo.Context) error {
// 	c.Echo().Logger.Debug("Booking")

// 	// get request
// 	tokenDetail := c.Get("myuser").(models.UserTokenDetails)
// 	user, err := u.userDB.FindOne(bson.M{"username": tokenDetail.Username})
// 	if err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	if len(user.Bookings) == 0 {
// 		return c.JSON(http.StatusNotFound, "no bookings")
// 	}

// 	bl := make([]interface{}, len(user.Bookings))
// 	for i, b := range user.Bookings {
// 		oid, err := primitive.ObjectIDFromHex(string(b))

// 		if err != nil {
// 			return err
// 		}
// 		bl[i] = bson.M{"_id": oid}
// 	}
// 	bookingHotels, err := u.hotelDB.Find(bson.M{"$or": bl})

// 	if err != nil {
// 		return err
// 	}
// 	//TODO: show more booking
// 	return c.JSON(http.StatusOK, bookingHotels)
// }

// // RegisterUser handlers
// func (u *UserHandler) RegisterUser(c echo.Context) error {
// 	c.Echo().Logger.Debug("Register")

// 	regisUser := c.Get("user").(models.User)

// 	//check email and username is already exist
// 	findUser, err := u.db.UserDB.FindOne(
// 		bson.M{"$or": []interface{}{
// 			bson.M{"username": regisUser.Username},
// 			bson.M{"email": regisUser.Email},
// 		}})
// 	c.Echo().Logger.Info("FIND")

// 	if findUser != nil {
// 		return c.JSON(http.StatusConflict, `{"message": already exist }`)
// 	}

// 	hash, err := u.b.Hash(regisUser.Password)
// 	c.Echo().Logger.Info("HASH")

// 	if err != nil {
// 		return err
// 	}
// 	// setting new password
// 	regisUser.ID = primitive.NewObjectID()
// 	regisUser.Password = hash

// 	c.Echo().Logger.Info("INSERT")

// 	err = u.userDB.Add(regisUser)
// 	c.Echo().Logger.Info("USER ADD", err)

// 	if err != nil {
// 		return err
// 	}
// 	c.Echo().Logger.Info("Insert User:", fmt.Sprintf("%+v\n", regisUser))
// 	return c.JSON(http.StatusOK, "success")
// }

// ListUsers handlers
// func (u *UserHandler) ListUsers(c echo.Context) error {

// 	//check email and username is already exist
// 	findUser, err := u.db.UserDB.

// 	if err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}
// 	if findUser == nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	return c.JSON(http.StatusOK, findUser)
// }

//TODO: delete

// LoginUser handlers is a login handlers to get models.login
// and compare password and hash return jwt token for store in front-side
func (u *UserHandler) LoginUser(c echo.Context) error {

	login := c.Get("login").(models.Login)

	// find user for login
	user, err := u.db.UserDB.FindUserByUsername(login.Username)

	if err == errors.ErrNoDocuments {
		return c.JSON(
			http.StatusNotFound,
			models.ErrorResponse{Error: errors.ErrNoDocuments.Error()},
		)
	}
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			models.ErrorResponse{Error: err.Error()},
		)
	}
	// not compare password
	if !u.b.Compare(login.Password, user.Password) {
		return c.JSON(
			http.StatusUnauthorized,
			models.ErrorResponse{Error: errors.ErrPasswordNotMatch.Error()},
		)
	}
	// create token
	token, err := u.jwt.CreateToken(user.Username)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			models.ErrorResponse{Error: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, models.TokenResponse{Token: fmt.Sprintf("Bearer %s", token)})
}
