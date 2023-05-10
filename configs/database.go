package configs

import (
	"task5-vix/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3307)/test?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}) 
	db.AutoMigrate(&models.Photo{})

	DB = db
	log.Println("Database connected")
}