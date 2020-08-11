package data

import (
	"errors"
	"subsea/models"
)



//TODO:
// booking

// HotelDB is an interface for interact with hotel Database
type HotelDB interface {
	CreateHotel(models.Hotel) (string, error)
	FindHotelByName(string) (*models.Hotel, error)
	ListAllHotels() (models.Hotels, error)
	RemoveHotelByName(string) (bool, error)
}

// UserDB is an interface for interact with users Database
type UserDB interface {
	CreateUser(models.User) (string, error)
	// IsUserExist given by email and username <- easy version
	// FindIsUserExists(string, string) (models.User, error)
	FindUserByUsername(string) (models.User, error)
	ListAllUsers() (models.Users, error)
}

// Database is a main database for application
type Database struct {
	UserDB  UserDB
	HotelDB HotelDB
}

// NewDatabase is constructor given by mongo client and name of db to create
func NewDatabase(userDB UserDB, hotelDB HotelDB) *Database {
	return &Database{UserDB: userDB, HotelDB: hotelDB}
}

// ShowUserBooking is a method for database using two database for booking
func (db *Database) ShowUserBooking(username string) (models.Hotels, error) {
	user, err := db.UserDB.FindUserByUsername(username)

	// check if found user
	if err == db.UserDB.
}
