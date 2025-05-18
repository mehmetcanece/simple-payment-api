package models

import "gorm.io/gorm"	


type Order struct {
	gorm.Model
	TableID uint
	TotalAmount float64
}