package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"Username"`
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
