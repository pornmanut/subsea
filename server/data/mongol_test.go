package data

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func connect(T *testing.T) *mongo.Client {
	client, err := ConnectMongoServer(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		T.Error(err)
	}
	return client
}

func TestAdd(T *testing.T) {
	client := connect(T)
	defer client.Disconnect(context.TODO())
	db, err := NewHotelMongo(client)

	if err != nil {
		T.Error(err)
	}
	h := Hotel{
		Name: "WAHA",
	}
	db.Add(h)
}

func TestList(T *testing.T) {
	client := connect(T)
	defer client.Disconnect(context.TODO())
	db, err := NewHotelMongo(client)

	if err != nil {
		T.Error(err)
	}

	result, err := db.List()

	if err != nil {
		T.Error(err)
	}
	fmt.Println(result)
}

func TestGet(T *testing.T) {
	client := connect(T)
	defer client.Disconnect(context.TODO())
	db, err := NewHotelMongo(client)

	if err != nil {
		T.Error(err)
	}

	result, err := db.Get("God")

	if err != nil {
		T.Error(err)
	}
	fmt.Println(result)
}

func TestDelete(T *testing.T) {
	client := connect(T)
	defer client.Disconnect(context.TODO())
	db, err := NewHotelMongo(client)

	if err != nil {
		T.Error(err)
	}

	err = db.Delete("God")

	if err != nil {
		T.Error(err)
	}
}

func TestCycle(T *testing.T) {
	client := connect(T)
	defer client.Disconnect(context.TODO())
	db, err := NewHotelMongo(client)

	if err != nil {
		T.Error(err)
	}
	h := Hotel{
		Name:   "Awesome",
		Price:  5000.01,
		Height: -300.34,
		Detail: "Nothing at all",
	}
	err = db.Add(h)

	if err != nil {
		T.Error(err)
	}

	result, err := db.Get("Awesome")

	if err != nil {
		T.Error(err)
	}

	if len(result) == 0 {
		T.Error(errors.New("Not found"))
	}

	err = db.Delete("Awesome")

	if err != nil {
		T.Error(err)
	}

	result, err = db.Get("Awesome")

	if err != nil {
		T.Error(err)
	}

	if len(result) > 0 {
		T.Error(errors.New("Can't Delete"))
	}

}
