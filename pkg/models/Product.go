package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string     `json:"name"`
	Price      float64    `gorm:"default:0.0" json:"price"`
	CategoryID []Category `gorm:"foreignKey:ID" json:"categoryID"`
	VendorID   []Vendor   `gorm:"foreignKey:ID" json:"vendorID"`
}
