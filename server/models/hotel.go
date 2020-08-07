package models

// Hotels is a collection of hotel
type Hotels []Hotel

// Hotel collection contain basic information and avliable for user can booking
type Hotel struct {
	Name    string  `json:"name" validate:"required"`
	Price   float32 `json:"price" validate:"required"`
	Detail  string  `json:"detail" validate:"required"`
	Height  float32 `json:"height" validate:"required"`
	Booking int     `json:"booking"`
	Max     int     `json:"max"`
}
