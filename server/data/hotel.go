package data

import "subsea/models"

// HotelsDB represent a hotel interface
type HotelsDB interface {
	List() (models.Hotels, error)
	Add(models.Hotel) error
	Get(string) (models.Hotels, error)
	Delete(string) error
}
