package model

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Fullname string `gorm:"type:varchar(50)"`
	Email string `gorm:"type:varchar(100);unique_index"`
	PasswordHash string `gorm:"type:varchar(60)"`
}