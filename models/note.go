package models

import "gorm.io/gorm"

// Note is the model for the notes table
type Note struct {
	gorm.Model
	Title   string
	Content string 
	UserId	uint
	User		User `gorm:"foreignkey:UserId"`   
}