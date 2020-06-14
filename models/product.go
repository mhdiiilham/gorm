package model

import (
	"github.com/jinzhu/gorm"
)

// Product model
type Product struct {
	gorm.Model
	Name string `gorm:"type:varchar(50)"`
	Price int `gorm:"type:int(30)"`
}
