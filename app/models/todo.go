package models

import (
	"github.com/jinzhu/gorm"
)

// Todo models
type Todo struct {
	gorm.Model
	Name string
	Role string `gorm:"size:255"` // set field size to 255
}
