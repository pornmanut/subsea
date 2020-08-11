package data

import (
	"context"
	"fmt"
	"log"
	"subsea/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO:
// func methods FindHotelByID
// func methods DeleteHotelByID
// func method SerachHotel with params <- Dificult. Easy way just get 4 parameter

// HotelMongoDB represent a hotels struct
type HotelMongoDB struct {
	collection *mongo.Collection
	log        *log.Logger
}

// NewHotelDB defined by constructor methods
func NewHotelDB(db *mongo.Database) *HotelMongoDB {
	col := db.Collection("hotels")

	return &HotelMongoDB{collection: col}
}

// CreateHotel is a method for create new hotel given by hotel model
// return with string of id and error from insert into mongoDB collection
func (db *HotelMongoDB) CreateHotel(hotel models.Hotel) (string, error) {
	// setting up context time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// setting up new Object ID
	hotel.ID = primitive.NewObjectID()

	// insert hotel into mongoDB database.
	result, err := db.collection.InsertOne(ctx, hotel)

	// checking error
	if err != nil {
		return "", err
	}

	// using InsertOneResult interface covert to ObjectID
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", ErrObjectID
	}
	// return into hex
	return id.Hex(), nil
}

// FindHotelByName is a methods for find one hotel given by name
// return with hotel models and error from mongoDB collection
func (db *HotelMongoDB) FindHotelByName(name string) (*models.Hotel, error) {
	// setting up context time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// find one hotel from mongoDB
	res := db.collection.FindOne(ctx, bson.M{name: name})

	// handling error from case not found
	err := res.Err()
	if err != nil {
		return nil, err
	}
	// create new Hotel
	hotel := new(models.Hotel)
	// decode result into hotel struct
	err = res.Decode(hotel)

	if err != nil {
		return nil, err
	}

	return hotel, nil
}

// ListAllHotels is a method for list all hotel
// return with a collection of hotel with error from mogoDB
func (db *HotelMongoDB) ListAllHotels() (models.Hotels, error) {
	// setting up context time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// findhotel from mongoDB
	cursor, err := db.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	// create new hotels
	var hotels models.Hotels

	// cursor.Next is a key to find another hotel
	// return true if found.
	// return false when ctx timeout or not found
	for cursor.Next(ctx) {
		// create single hotel
		hotel := new(models.Hotel)
		err := cursor.Decode(hotel)

		if err != nil {
			return nil, err
		}

		// append with copy value from hotel
		hotels = append(hotels, *hotel)
	}

	return hotels, nil
}

// RemoveHotelByName is a methods for remove one hotel given by name
// return with hotel models and error from mongoDB collection
func (db *HotelMongoDB) RemoveHotelByName(name string) (bool, error) {
	// setting up context time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// remove hotel from mongoDB
	deleteResult, err := db.collection.DeleteOne(ctx, bson.M{name: name})

	// handling error
	if err != nil {
		return false, err
	}
	// not found any document for delete
	if deleteResult.DeletedCount == 0 {
		return false, nil
	}
	// success deleting
	return true, nil
}

//Add add one records give by User struct
func (db *HotelMongoDB) Add(item models.Hotel) error {
	result, err := db.collection.InsertOne(context.TODO(), item)
	fmt.Println(result)
	return err
}

//FindOne find only one records from collection returns to user
func (db *HotelMongoDB) FindOne(filter bson.M) (*models.Hotel, error) {
	cursor := db.collection.FindOne(context.TODO(), filter)
	var hotel models.Hotel
	err := cursor.Decode(&hotel)
	if err != nil {
		return nil, err
	}
	return &hotel, err
}

//Find find many records from collection given by filter returns to a collection of user
func (db *HotelMongoDB) Find(filter interface{}) (models.Hotels, error) {
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
func (db *HotelMongoDB) DeleteOne(filter bson.M) error {
	result, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

//ReplaceOne find many records from collection given by filter returns to a collection of user
func (db *HotelMongoDB) ReplaceOne(filter bson.M, hotel models.Hotel) error {
	result, err := db.collection.ReplaceOne(context.TODO(), filter, hotel)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}
