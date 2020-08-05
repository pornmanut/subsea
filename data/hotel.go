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

// Hotels is a collection of hotel
type Hotels []Hotel

// Hotel collection contain basic information and avliable for user can booking
type Hotel struct {
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Detail  string  `json:"detail"`
	Height  float32 `json:"height"`
	Booking int     `json:"booking"`
	Max     int     `json:"max"`
	Open    bool    `json:"open"`
}

// HotelsDB represent a hotel collection databases
type HotelsDB struct {
	cl *mongo.Collection
}

// NewHotelDB create a new database
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
	collection := client.Database("subsea").Collection("hotels")

	return &HotelsDB{cl: collection}
}

func (db *HotelsDB) AddHotel(h Hotel) {
	result, err := db.cl.InsertOne(context.TODO(), h)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func (db *HotelsDB) GetHotels() Hotels {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cursor, err := db.cl.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
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
	return result
}
