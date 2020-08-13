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

// RegisterUser handler is a register handler for put model.User to database
// also validate exists user
func (u *UserHandler) RegisterUser(c echo.Context) error {
	u.l.Info("Register User Handler")

	user := c.Get("user").(models.User)

	fmt.Println(user)
	if err := u.isEmailExists(user.Email); err != nil {
		return c.JSON(err.Code, err)
	}

	if err := u.isUsernameExists(user.Username); err != nil {
		return c.JSON(err.Code, err)
	}

	// hashing password for store
	hash, err := u.b.Hash(user.Password)

	if err := handleUnknowError(err); err != nil {
		return c.JSON(err.Code, err)
	}

	// set new password
	user.Password = hash

	id, err := u.db.UserDB.CreateUser(user)

	if err := handleUnknowError(err); err != nil {
		return c.JSON(err.Code, err)
	}

	u.l.Info("Succesfuly create new user", "ID", id)
	return c.JSON(
		http.StatusOK,
		models.SuccessCreated{ID: id, Message: "Sucessfully create new user"},
	)
}

// LoginUser handlers is a login handlers to get models.login
// and compare password and hash return jwt token for store in front-side
func (u *UserHandler) LoginUser(c echo.Context) error {

	login := c.Get("login").(models.Login)

	// find user for login
	user, err := u.db.UserDB.FindUserByUsername(login.Username)

	if user == nil {
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
	// not compare password
	if !u.b.Compare(login.Password, user.Password) {
		return c.JSON(
			http.StatusUnauthorized,
			models.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: errors.ErrPasswordNotMatch.Error(),
			},
		)
	}
	// create token
	token, err := u.jwt.CreateToken(user.Username)

	if err := handleUnknowError(err); err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(
		http.StatusOK,
		models.TokenResponse{Token: fmt.Sprintf("Bearer %s", token)},
	)
}

// unexport

func handleUnknowError(err error) *models.ErrorResponse {
	if err != nil {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

func (u *UserHandler) isEmailExists(email string) *models.ErrorResponse {
	u.l.Info("Checking Email", "email", email)
	user := new(models.User)
	user, err := u.db.UserDB.FindUserByEmail(email)
	if err := handleUnknowError(err); err != nil {
		return err
	}

	// user exists
	if user != nil {
		u.l.Error("Conflict", "email", email)
		return &models.ErrorResponse{
			Code:    http.StatusConflict,
			Message: errors.ErrEmailAlreadyExists.Error(),
		}
	}
	return nil
}

func (u *UserHandler) isUsernameExists(username string) *models.ErrorResponse {
	u.l.Info("Checking Email", "username", username)

	user := new(models.User)
	user, err := u.db.UserDB.FindUserByUsername(username)

	if err := handleUnknowError(err); err != nil {
		return err
	}
	// user exists
	if user != nil {
		u.l.Error("Conflict", "username", username)
		return &models.ErrorResponse{
			Code:    http.StatusConflict,
			Message: errors.ErrUsernameAlreadyExists.Error(),
		}
	}
	return nil
}
