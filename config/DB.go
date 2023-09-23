package config

import (
	"models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DB_init() {
	database, err := gorm.Open(sqlite.Open("DB.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	err = database.AutoMigrate(&models.Player{})
	if err != nil {
		return
	}

	DB = database
}
