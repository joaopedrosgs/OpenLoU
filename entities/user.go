package entities

import (
	"time"
)

type User struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string
	Email        string
	PasswordHash string
}
