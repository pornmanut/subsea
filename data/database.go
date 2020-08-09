package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectMongoServer connect to mongo server with url
func ConnectMongoServer(ctx context.Context, url string) (*mongo.Client, error) {
	// Configure a Client with SCRAM authentication (https://docs.mongodb.com/manual/core/security-scram/).
	// The default authentication database for SCRAM is "admin". This can be configured via the
	// authSource query parameter in the URI or the AuthSource field in the options.Credential struct.
	// SCRAM is the default auth mechanism so specifying a mechanism is not required.

	// To configure auth via URI instead of a Credential, use
	// "mongodb://user:password@localhost:27017".
	// credential := options.Credential{
	// 	Username:      username,
	// 	Password:      password,
	// 	AuthMechanism: "SCRAM-SHA-1",
	// }
	// fmt.Println(credential)
	// .SetAuth(credential)
	clientOptions := options.Client().ApplyURI(url + "?retryWrites=false")
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

// Database define for database for this application
type Database struct {
	UserDB  *UserDB
	HotelDB *HotelDB
}

// NewDatabase is constructor given by mongo client and name of db to create
func NewDatabase(client *mongo.Client, nameOfDB string) *Database {

	db := client.Database("subsea")
	userDB := NewUserDB(db)
	hotelDB := NewHotelDB(db)

	return &Database{UserDB: userDB, HotelDB: hotelDB}
}
