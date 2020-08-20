package mongo

import (
	"context"
	"fmt"
	"subsea/models"
	"testing"
	"time"
)

func setupMongoDB(T *testing.T) *UserMongoDB {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := ConnectMongoServer(ctx, "mongodb://localhost:27017")

	if err != nil {
		T.Error(err)
	}

	db := client.Database("test")
	userDB := NewUserDB(db)
	return userDB
}

func TestCreateDelete(T *testing.T) {
	db := setupMongoDB(T)

	user := models.User{
		Email:    "god@gmail.com",
		Username: "god",
		Password: "test",
	}
	id, err := db.CreateUser(user)

	if err != nil {
		T.Error(err)
	}

	err = db.RemoveUserByUserName(user.Username)

	if err != nil {
		T.Error(err)
	}
	fmt.Println(id)
}
