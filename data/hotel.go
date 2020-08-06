package data

// Hotels is a collection of hotel
type Hotels []Hotel

// Hotel collection contain basic information and avliable for user can booking
type Hotel struct {
	ID      string  `json:"id" bson:"_id"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Detail  string  `json:"detail"`
	Height  float32 `json:"height"`
	Booking int     `json:"booking"`
	Max     int     `json:"max"`
	Open    bool    `json:"open"`
}

// HotelsDB represent a hotel interface
type HotelsDB interface {
	List() (Hotels, error)
	Add(Hotel) error
	Get(string) (Hotels, error)
	Delete(string) error
}
