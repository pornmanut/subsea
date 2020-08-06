package data

import (
	"context"
	"fmt"
	"log"

	"subsea/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserDB is user database collection mongoCollection
type UserDB struct {
	collection *mongo.Collection
}

// NewUserDB following by constructor methods
func NewUserDB(client *mongo.Client) *UserDB {
	col := client.Database("subsea").Collection("users")
	return &UserDB{collection: col}
}

//Add add one records give by User struct
func (db *UserDB) Add(user models.User) error {
	result, err := db.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

//FindOne find only one records from collection returns to user
func (db *UserDB) FindOne(filter bson.M) (*models.User, error) {
	cursor := db.collection.FindOne(context.TODO(), filter)
	var user models.User
	err := cursor.Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, err
}

//Find find many records from collection given by filter returns to a collection of user
func (db *UserDB) Find(filter bson.M) (models.Users, error) {
	cursor, err := db.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var result models.Users
	for cursor.Next(context.TODO()) {
		var user models.User
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		result = append(result, user)
	}
	return result, nil
}

// DeleteOne delete one recrods from collections given by filter
func (db *UserDB) DeleteOne(filter bson.M) error {
	result, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

// func (db *UserDB)
