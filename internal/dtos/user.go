package dtos

import "github.com/alihaqberdi/goga_go/internal/models/types"

type User struct {
	ID       uint           `json:"id"`
	Username string         `json:"username"`
	Role     types.UserRole `json:"role"`
	Email    string         `json:"email"`
}
