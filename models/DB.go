package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("DB.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
	}

	err = db.AutoMigrate(&Quest{})
	if err != nil {
		log.Println(err)
	}

	err = db.AutoMigrate(&Character{})
	if err != nil {
		log.Println(err)
	}

	return db, err
}
