package models

import "gorm.io/gorm"

type Quest struct {
	gorm.Model
	Name       string `json:"Name"`
	Difficulty int    `json:"Difficulty"`
}
