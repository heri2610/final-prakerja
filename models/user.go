package models

import (
	"time"
)

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
