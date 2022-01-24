package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	ID             int
	Name           string
	Characteristic string
	Value          string
}
