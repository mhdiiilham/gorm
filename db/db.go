package db

import (
	"github.com/jinzhu/gorm"
	// 
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"

	"github.com/mhdiiilham/gorm/models"
)

// Connection ...
func Connection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/company?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&model.Product{})
	if err != nil {log.Fatal(err)}
	log.Info("Database connected")
	return db
}