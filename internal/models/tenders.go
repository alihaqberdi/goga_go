package models

import (
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"time"
)

type Tender struct {
	Id          int					`gorm:"primaryKey";autoIncrement`
	ClientId    int					`gorm:"not null"`
	Title       string 				`gorm:"type:varchar(255);not null"`
	Description string 				`gorm:"type:text"`
	Deadline    time.Time			`gorm:"not null"`
	Budget      float64 			`gorm:"not null"`
	Status      types.TenderStatus 	`gorm:"not null"`
}
