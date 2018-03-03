package entities

import (
	"time"
)

type User struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string
	Email        string
	PasswordHash string
	AllianceName string
	Gold         uint
	Diamonds     uint
	Darkwood     uint
	Runestone    uint
	Veritium     uint
	Trueseed     uint
	Rank         uint
}
