package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string 
	Hash     []byte
	Email    string `gorm:"not null;unique"`
	Notes		 []Note
}
