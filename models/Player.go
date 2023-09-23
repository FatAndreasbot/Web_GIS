package models

import "gorm.io/gorm"

type Player struct {
	Name string
	gorm.Model
}
