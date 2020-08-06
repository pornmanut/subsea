package handlers

import (
	"net/http"
	"subsea/data"

	"github.com/labstack/echo/v4"
)

// MiddlewareValidateUser validates the user in the request and call net it ok
func (u *UserHandler) MiddlewareValidateUser(next echo.HandlerFunc) echo.HandlerFunc {
	// header
	return func(c echo.Context) error {

		var user data.User

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
		// call next
		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}

}
