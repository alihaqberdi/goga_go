package dtos

import "github.com/alihaqberdi/goga_go/internal/models/types"

type JwtUser struct {
	Id       int            `json:"id"`
	Username string         `json:"username"`
	Role     types.UserRole `json:"role"`
}
