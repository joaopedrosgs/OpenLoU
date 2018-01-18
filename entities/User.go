package entities

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index"`
	Email string `gorm:"type:varchar(100);unique_index"`
	PasswordHash string
	Cities []City
}