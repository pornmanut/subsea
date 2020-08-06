package models

// Booking define a string match to object_id of hotels collection
type Booking string

// User struct is an informaton of signle user and booking
type User struct {
	Email     string    `json:"email" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	FirstName string    `json:"fistname"`
	LastName  string    `json:"lastname"`
	BirthDate string    `json:"brithdate"`
	Bookings  []Booking `json:"bookings"`
}

// Users is a collction of user
type Users []User
