package models

import (
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string         `gorm:"type:varchar(255);not null";unique`
	PasswordHash string         `gorm:"not null"`
	Role         types.UserRole `gorm:"not null"`
	Email        string         `gorm:"type:varchar(255);not null";unique`
}
