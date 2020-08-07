package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongoServer(ctx context.Context, url string) (*mongo.Client, error) {
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

type Database struct {
	UserDB  *UserDB
	HotelDB *HotelDB
}

func NewDatabase(client *mongo.Client, nameOfDB string) *Database {
	db := client.Database("subsea")

	userDB := NewUserDB(db)
	hotelDB := NewHotelDB(db)

	return &Database{UserDB: userDB, HotelDB: hotelDB}
}
