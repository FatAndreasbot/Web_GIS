package config

import (
	"log"
	"models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

var PrivateKey = []byte(os.Getenv("API_SECRET"))

func DB_init() {
	database, err := gorm.Open(sqlite.Open("DB.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	err = database.AutoMigrate(&models.Player{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	DB = database
}

func Env_init() {
	os.Setenv("TOKEN_HOUR_LIFESPAN", "8")
	os.Setenv("API_SECRET", "dc0fa9d3c5673a432767928515739131c120bad50d7009c21e88fd9ea9e8d976")
}
