package models

import "gorm.io/gorm"

type Vendor struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
}
