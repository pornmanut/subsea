package data

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestUserGet(T *testing.T) {
	client := connect(T)
	defer client.Disconnect(context.TODO())
	db := NewUserDB(client)

	a := db.Get("A")

	b, err := json.Marshal(a)
	if err != nil {
		T.Error(err)
	}
	fmt.Println(string(b))
}
