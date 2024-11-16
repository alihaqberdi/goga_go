package models

import (
	"time"

	"github.com/alihaqberdi/goga_go/internal/models/types"
	"gorm.io/gorm"
)

type Bid struct {
	gorm.Model
	TenderId     int             `gorm:"not null"`
	ContractorId int             `gorm:"not null"`
	Price        float64         `gorm:"not null"`
	DeliveryTime time.Time       `gorm:"not null"`
	Comments     string          `gorm:"type:text"`
	Status       types.BidStatus `gorm:"not null"`
}
