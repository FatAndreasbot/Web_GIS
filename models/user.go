package models

import (
	"errors"
	"html"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string `gorm:"size:255;not null;unique" json:"Login"`
	Username string `gorm:"size:255;not null" json:"Username"`
	Password string `gorm:"size:255;not null;" json:"-"`
}

func (u *User) HashPassword() error {
	hashedPWord, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPWord)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func GetUserByID(userID uint) (User, error) {
	var user User

	db, err := SetupDB()
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}

	err = db.Table("users").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}
