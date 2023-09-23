package models

import "gorm.io/gorm"

type Player struct {
	Name      string
	ModelData gorm.Model
	// gorm.Model
}
