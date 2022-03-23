package user

import "time"

// User struct
type User struct {
	ID             int
	Name           string
	Gender         string
	Email		   string
	PhoneNumber    string
	Alamat         string
	PasswordHash   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
