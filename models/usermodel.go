package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Username  string
	Email     string `gorm:"unique;not null"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Photos    []Photo `gorm:"foreigKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Register struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MyProfile struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateProfile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
