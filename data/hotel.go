package data

// Hotels is a collection of hotel
type Hotels []Hotel

// Hotel collection contain basic information and avliable for user can booking
type Hotel struct {
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Detail  string  `json:"detail"`
	Height  float32 `json:"height"`
	Booking int     `json:"booking"`
	Max     int     `json:"max"`
	Open    bool    `json:"open"`
}

// HotelsDB represent a hotel collection databases
type HotelsDB interface {
	List() (Hotels, error)
	Add(Hotel) error
}
