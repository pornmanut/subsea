package errors

import "errors"

// ErrNoDocuments is an error. when cannot find any result
var ErrNoDocuments = errors.New("Not found any document on request")

// ErrHotelAlreadyExists an error. when conflict with hotel name
var ErrHotelAlreadyExists = errors.New("Hotel already exists")

// ErrPasswordNotMatch an error. when login and password doesn't match
var ErrPasswordNotMatch = errors.New("Password not match")

// ErrEmailAlreadyExists an error. when email is exists
var ErrEmailAlreadyExists = errors.New("Email already exists")

// ErrUsernameAlreadyExists an error. when username is exists
var ErrUsernameAlreadyExists = errors.New("Username already exists")
