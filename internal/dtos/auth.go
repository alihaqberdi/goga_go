package dtos

import "github.com/alihaqberdi/goga_go/internal/models/types"

type Register struct {
	Username string         `json:"username"`
	Password string         `json:"password"`
	Role     types.UserRole `json:"role"`
	Email    string         `json:"email"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRes struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
