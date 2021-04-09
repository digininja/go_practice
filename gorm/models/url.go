// models/book.go

package models

import (
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	URL    string `json:"url"`
	Status string `json:"status"`
}
