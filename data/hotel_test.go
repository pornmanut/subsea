package data

import (
	"fmt"
	"testing"
)

func TestAdd(T *testing.T) {
	db := NewHotelDB()

	h := Hotel{
		Name: "WAHA",
	}
	db.AddHotel(h)
}

func TestList(T *testing.T) {
	db := NewHotelDB()

	result := db.GetHotels()

	fmt.Println(result)
	T.Error()
}
