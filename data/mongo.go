package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	col *mongo.Collection
}

// ConnectMongoDB connect to mongoDB database
func ConnectMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewHotelDB create a new database
func NewHotelMongo() (*MongoDB, error) {
	client, err := ConnectMongoDB()

	if err != nil {
		return nil, err
	}

	col := client.Database("subsea").Collection("hotels")
	return &MongoDB{col: col}, nil
}

func (db *MongoDB) Add(h Hotel) error {
	result, err := db.col.InsertOne(context.TODO(), h)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func (db *MongoDB) List() (Hotels, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cursor, err := db.col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result Hotels
	for cursor.Next(ctx) {
		var hotel Hotel
		if err = cursor.Decode(&hotel); err != nil {
			log.Fatal(err)
		}
		result = append(result, hotel)
	}
	return result, nil
}

func (db *MongoDB) Get(name string) (Hotels, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cursor, err := db.col.Find(ctx, bson.M{"name": name})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result Hotels
	for cursor.Next(ctx) {
		var hotel Hotel
		if err = cursor.Decode(&hotel); err != nil {
			log.Fatal(err)
		}
		result = append(result, hotel)
	}
	return result, nil
}
