package entities

import (
	"time"
)

type User struct {
	ID           uint `gorm:"primary_key"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string `gorm:"type:varchar(100);unique_index"`
	Email        string `gorm:"type:varchar(100);unique_index"`
	PasswordHash string
	Cities       []City `gorm:"ForeignKey:UserID"`
}
