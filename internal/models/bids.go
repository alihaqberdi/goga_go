package models

import (
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"time"
)

type Bid struct {
	Id           int
	TenderId     int
	ContractorId int
	Price        float64
	DeliveryTime time.Time
	Comments     string
	Status       types.BidStatus
}
