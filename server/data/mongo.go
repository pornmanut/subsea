package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	col *mongo.Collection
}

// ConnectMongoServer connect to mongoDB database
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

// NewHotelMongo create a new database
func NewHotelMongo(client *mongo.Client) (*MongoDB, error) {
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

func (db *MongoDB) Delete(name string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := db.col.DeleteOne(ctx, bson.M{"name": name})

	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
