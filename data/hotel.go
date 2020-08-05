package data

import "github.com/hashicorp/go-hclog"

type HotelsDB struct {
}

func NewHotelDB(l hclog.Logger) *HotelsDB {
	return &HotelsDB{}
}
