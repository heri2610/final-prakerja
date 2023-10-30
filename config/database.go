package config

import (
	"os"

	"github.com/heri2610/final-prakerja/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dns := os.Getenv("DNS")
	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&models.User{})
}
