package data

import (
	"context"
	"fmt"
	"log"
	"subsea/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//TODO: interface for local or not mongodb

// HotelDB represent a hotels struct
type HotelDB struct {
	collection *mongo.Collection
}

// NewHotelDB defined by constructor methods
func NewHotelDB(db *mongo.Database) *HotelDB {
	col := db.Collection("hotels")

	return &HotelDB{collection: col}
}

//Add add one records give by User struct
func (db *HotelDB) Add(item models.Hotel) error {
	result, err := db.collection.InsertOne(context.TODO(), item)
	fmt.Println(result)
	return err
}

//FindOne find only one records from collection returns to user
func (db *HotelDB) FindOne(filter bson.M) (*models.Hotel, error) {
	cursor := db.collection.FindOne(context.TODO(), filter)
	var hotel models.Hotel
	err := cursor.Decode(&hotel)
	if err != nil {
		return nil, err
	}
	return &hotel, err
}

//Find find many records from collection given by filter returns to a collection of user
func (db *HotelDB) Find(filter interface{}) (models.Hotels, error) {
	cursor, err := db.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var result models.Hotels
	for cursor.Next(context.TODO()) {
		var hotel models.Hotel
		if err = cursor.Decode(&hotel); err != nil {
			log.Fatal(err)
		}
		result = append(result, hotel)
	}
	return result, nil
}

// DeleteOne delete one recrods from collections given by filter
func (db *HotelDB) DeleteOne(filter bson.M) error {
	result, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

//ReplaceOne find many records from collection given by filter returns to a collection of user
func (db *HotelDB) ReplaceOne(filter bson.M, hotel models.Hotel) error {
	result, err := db.collection.ReplaceOne(context.TODO(), filter, hotel)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}
