package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int `gorm:"serial"`
	FName    string
	UName    string
	Lame     string
	Password string
	Email    string
}
