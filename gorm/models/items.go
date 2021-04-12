// models/book.go

package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name    string
	Price   int
	OrderID uint
}
