package models

import (
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"gorm.io/gorm"
)

type Bid struct {
	gorm.Model
	TenderId     uint            `gorm:"not null"`
	ContractorId uint            `gorm:"not null"`
	Price        float64         `gorm:"not null"`
	DeliveryTime uint            `gorm:"not null"`
	Comments     string          `gorm:"type:text"`
	Status       types.BidStatus `gorm:"not null"`
}
