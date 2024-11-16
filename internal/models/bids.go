package models

import (
	"time"

	"github.com/alihaqberdi/goga_go/internal/models/types"
	"gorm.io/gorm"
)

type Bid struct {
	gorm.Model
	TenderId     int             `gorm:"not null" json:"tender_id"`
	ContractorId int             `gorm:"not null" json:"contractor_id"`
	Price        float64         `gorm:"not null" json:"price"`
	DeliveryTime time.Time       `gorm:"not null" json:"delivery_time"`
	Comments     string          `gorm:"type:text" json:"comments"`
	Status       types.BidStatus `gorm:"not null" json:"status"`
}
