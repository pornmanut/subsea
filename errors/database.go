package errors

import "errors"

// ErrNoDocuments is an error. when cannot find any result
var ErrNoDocuments = errors.New("Not found any document on request")

// ErrHotelAlreadyExists an error. when conflict with hotel name
var ErrHotelAlreadyExists = errors.New("Hotel already exists")

// ErrPasswordNotMatch an error. when login and password doesn't match
var ErrPasswordNotMatch = errors.New("Password not match")
