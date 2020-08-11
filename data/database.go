package data

import (
	"context"
	"errors"
	"subsea/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ErrObjectID is an error. when cannot covert objectID
var ErrObjectID = errors.New("Can not covert object into primitive ObjectID")

// ErrNoDocuments is an error. when cannot find any result
var ErrNoDocuments = errors.New("Not found any document on request")

//TODO:
// booking

// ConnectMongoServer connect to mongo server with url
func ConnectMongoServer(ctx context.Context, url string) (*mongo.Client, error) {

	// url = url + "?retryWrites=false"
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, err
	}

	return client, nil
}

// HotelDB is an interface for interact with hotel Database
type HotelDB interface {
	CreateHotel(models.Hotel) (string, error)
	FindHotelByName(string) (models.Hotel, error)
	ListAllHotels() (models.Hotels, error)
	RemoveHotelByName(string) (bool, error)
}

// UserDB is an interface for interact with users Database
type UserDB interface {
	CreateUser(models.User) (string, error)
	// IsUserExist given by email and username <- easy version
	IsUserExist(string, string) (models.User, error)
	ListAllUsers() (models.Users, error)
}

// Database is a main database for application
type Database struct {
	UserDB  *UserMongoDB
	HotelDB *HotelMongoDB
}

// NewDatabase is constructor given by mongo client and name of db to create
func NewDatabase(client *mongo.Client, nameOfDB string) *Database {

	db := client.Database(nameOfDB)
	userDB := NewUserDB(db)
	hotelDB := NewHotelDB(db)

	return &Database{UserDB: userDB, HotelDB: hotelDB}
}

// ShowBooking is a method for database using two database for booking
func (db *Database) ShowBooking(username string) {
}
