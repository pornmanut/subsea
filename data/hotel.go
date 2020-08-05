package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HotelsDB struct {
	cl *mongo.Collection
}

func NewHotelDB() *HotelsDB {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}
	collection := client.Database("mydb").Collection("persons")

	return &HotelsDB{cl: collection}
}
