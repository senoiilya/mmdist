package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID []User `gorm:"foreignKey:ID" json:"userId"`
}
