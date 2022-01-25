package models

import "gorm.io/gorm"

type PropertyCharacteristic struct {
	gorm.Model
	PropertyId     int
	Characteristic string
	Value          string
}
