package models

import (
	"time"
	"github.com/jackc/pgx"
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

func GetUserByName(db pgx.Conn, name string) *User {

}
func GetUserByEmail(db pgx.Conn, email string) *User {

}
