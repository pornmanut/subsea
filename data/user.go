package data

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	FirstName string    `json:"fistname"`
	LastName  string    `json:"lastname"`
	BirthDate string    `json:"brithdate"`
	Bookings  []Booking `json:"bookings"`

	password string
}

type Booking string

type UserDB struct {
	client  *mongo.Client
	userCol *mongo.Collection
}

func NewUserDB(client *mongo.Client) *UserDB {
	col := client.Database("subsea").Collection("users")
	return &UserDB{client: client, userCol: col}
}

func (db *UserDB) Get(email string) *User {
	return &User{
		Email:     "god",
		Username:  "yea",
		FirstName: "wod",
		LastName:  "haha",
		BirthDate: "todo",
		Bookings: []Booking{
			"abc",
			"dec",
		},
		password: "123",
	}
}

// func (db *UserDB)
