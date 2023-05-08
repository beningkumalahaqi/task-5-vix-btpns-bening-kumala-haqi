package configs

import (
	"log"
	"task5-vix/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=vix_db port=5432 sslmode=disable TimeZone=Asia/Singapore"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})

	DB = db
	log.Println("Database connected")
}
