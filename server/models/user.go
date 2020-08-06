package models

// Booking define a string match to object_id of hotels collection
type Booking string

// User struct is an informaton of signle user and booking
type User struct {
	Email     string    `json:"email" validate:"required,email"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	FirstName string    `json:"firstname" validate:"required"`
	LastName  string    `json:"lastname"  validate:"required"`
	BirthDate string    `json:"birthdate" validate:"required,date"`
	Bookings  []Booking `json:"bookings"`
}

// Users is a collction of user
type Users []User
