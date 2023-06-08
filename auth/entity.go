package auth

import "time"

type User struct {
	ID        int
	FullName  string
	UserName  string
	Email     string `gorm:"unique"`
	Password  string
	Phone     string
	Type      string
	Address   string
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
