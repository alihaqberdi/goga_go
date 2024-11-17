package types

type UserRole string

func (r UserRole) Valid() bool {
	return validUserRole[r]
}

const (
	UserRoleClient     UserRole = "client"
	UserRoleContractor UserRole = "contractor"
)

var validUserRole = map[UserRole]bool{
	UserRoleClient:     true,
	UserRoleContractor: true,
}
