package data

import (
	"fmt"
	"testing"
)

func TestAdd(T *testing.T) {
	db, err := NewHotelMongo()

	if err != nil {
		T.Error(err)
	}
	h := Hotel{
		Name: "WAHA",
	}
	db.Add(h)
}

func TestList(T *testing.T) {
	db, err := NewHotelMongo()

	if err != nil {
		T.Error(err)
	}

	result, err := db.List()

	if err != nil {
		T.Error(err)
	}
	fmt.Println(result)
}
