package dtos

import (
	"time"

	"github.com/alihaqberdi/goga_go/internal/models/types"
)

type BidsBase struct {
	TenderID     uint            `json:"tender_id" example:"1"`
	ContractorID uint            `json:"contractor_id" example:"1"`
	Price        float64         `json:"price" example:"100.00"`
	DeliveryTime time.Time       `json:"delivery_time" example:"12"`
	Comments     string          `json:"comments" example:"This is a comment"`
	Status       types.BidStatus `json:"status" example:"pending"`
}

type BidCreate struct {
	BidsBase
}

type BidList struct {
	BidsBase
	ID uint `json:"id"`
}
