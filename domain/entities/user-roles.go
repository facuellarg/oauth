package entities

type UserRoleEnum string

const (
	admin UserRoleEnum = "admin"
	user  UserRoleEnum = "user"
)

type UserRole struct {
	Role UserRoleEnum `json:"role"`
}
