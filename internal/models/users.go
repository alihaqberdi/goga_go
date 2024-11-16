package models

import "github.com/alihaqberdi/goga_go/internal/models/types"

type User struct {
	Id       int
	Username string
	Password string
	Role     types.UserRole
	Email    string
}
