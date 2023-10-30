package config

import (
	"fmt"
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
		fmt.Println(err)
	}
	// Migration()
}

func Migration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.SosialMedia{})
}
