package mongo

import (
	"context"
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
func (db *UserMongoDB) FindUserByUsername(username string) (*models.User, error) {
	// setting up context time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// find one user from mongoDB
	res := db.collection.FindOne(ctx, bson.M{"username": username})

	// handling error from case not found
	err := res.Err()

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
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

// FindUserByEmail is a methods for find one User given by email
// return with username models and error from mongoDB collection
func (db *UserMongoDB) FindUserByEmail(email string) (*models.User, error) {
	// setting up context time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// find one user from mongoDB
	res := db.collection.FindOne(ctx, bson.M{"email": email})

	// handling error from case not found
	err := res.Err()

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

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
