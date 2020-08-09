package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Booking define a string match to object_id of hotels collection
type Booking string

// User struct is an informaton of signle user and booking
type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email     string             `json:"email" validate:"required,email"`
	Username  string             `json:"username" validate:"required"`
	Password  string             `json:"password" validate:"required"`
	FirstName string             `json:"firstname" validate:"required"`
	LastName  string             `json:"lastname"  validate:"required"`
	BirthDate string             `json:"birthdate" validate:"required,date"`
	Bookings  []Booking          `json:"bookings"`
}

// Users is a collction of user
type Users []User

//Login is login from to pass and validate
type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserTokenDetails struct {
	Username string `json:"username"`
}
