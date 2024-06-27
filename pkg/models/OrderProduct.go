package models

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model
	OrderID    []Order  `gorm:"foreignKey:ID" json:"orderID"`
	VendorID   []Vendor `gorm:"foreignKey:ID" json:"vendorID"`
	TotalPrice float64  `gorm:"default:0.0" json:"totalPrice"`
	Count      uint     `gorm:"default:0" json:"count"`
}
