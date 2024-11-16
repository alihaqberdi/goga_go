package dtos

import "github.com/alihaqberdi/goga_go/internal/models/types"

type Register struct {
	Username string
	Password string
	Role     types.UserRole
	Email    string
}

type Login struct {
	Username string
	Password string
	Role     types.UserRole
	Email    string
}
