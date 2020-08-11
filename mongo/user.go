package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"subsea/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserMongoDB using mongoDB
type UserMongoDB struct {
	collection *mongo.Collection
}

// NewUserDB following by constructor methods
func NewUserDB(db *mongo.Database) *UserMongoDB {
	col := db.Collection("users")
	return &UserMongoDB{collection: col}
}

// CreateUser is a method for create new user given by user model
// return with string of id and error from mongoDB collecton
func (db *UserMongoDB) CreateUser(user models.User) (string, error) {
	// setting up context time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// setting up new Object ID
	user.ID = primitive.NewObjectID()

	// insert user into mongoDB database.
	result, err := db.collection.InsertOne(ctx, user)

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

// FindUserByUsername is a methods for find one User given by username
// return with username models and error from mongoDB collection
func (db *HotelMongoDB) FindUserByUsername(username string) (*models.User, error) {
	// setting up context time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// find one user from mongoDB
	res := db.collection.FindOne(ctx, bson.M{username: username})

	// handling error from case not found
	err := res.Err()
	if err != nil {
		return nil, err
	}
	// create new Hotel
	user := new(models.User)
	// decode result into hotel struct
	err = res.Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

//Add add one records give by User struct
func (db *UserMongoDB) Add(item models.User) error {
	result, err := db.collection.InsertOne(context.TODO(), item)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

//FindOne find only one records from collection returns to user
func (db *UserMongoDB) FindOne(filter bson.M) (*models.User, error) {
	cursor := db.collection.FindOne(context.TODO(), filter)
	var user models.User
	err := cursor.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

//Find find many records from collection given by filter returns to a collection of user
func (db *UserMongoDB) Find(filter bson.M) (models.Users, error) {
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
func (db *UserMongoDB) DeleteOne(filter bson.M) error {
	result, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

//ReplaceOne find many records from collection given by filter returns to a collection of user
func (db *UserMongoDB) ReplaceOne(filter bson.M, user models.User) error {
	result, err := db.collection.ReplaceOne(context.TODO(), filter, user)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}
