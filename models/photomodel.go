package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Title    string
	Caption  string
	PhotoUrl string
	UserID   uint `gorm:"unique"`
}

type UploadPhoto struct {
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID uint `json:"user_id"`
}

type UpdatePhoto struct {
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}
