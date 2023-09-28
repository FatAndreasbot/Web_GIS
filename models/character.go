package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	UserID    uint `gorm:"ForeignKey:UserID; unique"`
	Name      string
	MaxHealth int
	Health    int
	AC        int
	Strenght  int
	Dexterity int
}
