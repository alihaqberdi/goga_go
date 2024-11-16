package models

import "github.com/alihaqberdi/goga_go/internal/models/types"

type User struct {
	Id       int 			`gorm:"primaryKey";autoIncrement`
	Username string			`gorm:"type:varchar(255);not null";unique`
	Password string			`gorm:"not null"`
	Role     types.UserRole	`gorm:"not null"`
	Email    string			`gorm:"type:varchar(255);not null";unique`
}
