// models/book.go

package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Name  string
	Items []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
