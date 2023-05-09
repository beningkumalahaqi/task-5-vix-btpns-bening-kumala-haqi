package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Title    string
	Caption  string
	PhotoUrl string
	UserID   int
}
