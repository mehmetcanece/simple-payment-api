package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Amount float64
	Method string
	Status string
}

// payment modelini tanımladım. dbdeki tabloya karşılık geliypr