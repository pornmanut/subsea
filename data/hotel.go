package data

import "github.com/hashicorp/go-hclog"

type HotelsDB struct {
	log hclog.Logger
}

func NewHotelDB(l hclog.Logger) *HotelsDB {
	return &HotelsDB{log: l}
}
