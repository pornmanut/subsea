package data

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

// func TestUserGet(T *testing.T) {
// 	client := connect(T)
// 	defer client.Disconnect(context.TODO())
// 	db := NewUserDB(client)

// 	a := db.Get("A")

// 	b, err := json.Marshal(a)
// 	if err != nil {
// 		T.Error(err)
// 	}
// 	fmt.Println(string(b))
// }

func TestAddUser(T *testing.T) {
	client := connect(T)
	defer client.Disconnect(context.TODO())
	db := NewUserDB(client)

	err := db.Add(User{
		Email:     "god",
		Username:  "yea",
		FirstName: "wod",
		LastName:  "haha",
		BirthDate: "todo",
		Bookings: []Booking{
			"abc",
			"dec",
		},
		Password: "123",
	})

	if err != nil {
		T.Error(err)
	}
}

func TestFindUser(T *testing.T) {
	client := connect(T)
	defer client.Disconnect(context.TODO())
	db := NewUserDB(client)

	a, _ := db.FindOne(bson.M{"email": "god"})
	fmt.Println(a)
}

func TestDeleteUser(T *testing.T) {
	client := connect(T)
	defer client.Disconnect(context.TODO())
	db := NewUserDB(client)

	a := db.DeleteOne(bson.M{"email": "god"})
	fmt.Println(a)
}
