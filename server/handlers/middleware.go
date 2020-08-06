package handlers

import (
	"net/http"
	"subsea/models"

	"github.com/labstack/echo/v4"
)

// UserContext is custom echo context
type UserContext struct {
	echo.Context
}

// MiddlewareValidateUser validates the user in the request and call net it ok
func (u *UserHandler) MiddlewareValidateUser(next echo.HandlerFunc) echo.HandlerFunc {
	// header
	return func(c echo.Context) error {
		var user models.User
		// bind user
		if err := c.Bind(&user); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		// validate the user
		errs := u.v.Validate(user)

		if len(errs) > 0 {
			return c.JSON(http.StatusUnprocessableEntity, errs.Errors())
		}

		// TODO: check dulipcate email

		// add user to context
		c.Set("user", user)
		// call next
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}

}

// MiddlewareValidateLogin validates the login in the request and call net it ok
func (u *UserHandler) MiddlewareValidateLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var login models.Login
		// bind user
		if err := c.Bind(&login); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		// validate the user
		errs := u.v.Validate(login)

		if len(errs) > 0 {
			return c.JSON(http.StatusUnprocessableEntity, errs.Errors())
		}

		// add user to context
		c.Set("login", login)
		// call next
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

// // MiddlewareFindByUserName given by user. find it and set to context
// func (u *UserHandler) MiddlewareCheckUserNameIsExist(db *data.UserDB, next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		requestUser := c.Get("user").(models.User)
// 		user, err := db.Find(bson.M{"username": requestUser.Username})

// 		if err != nil {
// 			return err
// 		}

// 		if user
// 		// call next
// 		if err := next(c); err != nil {
// 			c.Error(err)
// 		}
// 		return nil
// 	}
// }
