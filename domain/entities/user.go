package entities

type User struct {
	Name     string   `json:"name" gorm:"not null"`
	Image    string   `json:"image"`
	Email    string   `json:"email" gorm:"unique,not null"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}
