package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Photos []Photo `gorm:"foreigKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Photo struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	Title    string
	Caption  string
	PhotoUrl string
	UserID   uint
}
