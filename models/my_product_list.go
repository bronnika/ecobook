package models

import "time"

type MyProductList struct {
	ID         int       `gorm:"column:p_id" json:"id"`
	Title      string    `gorm:"column:p_title" json:"title"`
	Category   string    `gorm:"column:p_category" json:"category"`
	Photo      string    `gorm:"column:p_photo" json:"photo"`
	Price      string    `gorm:"column:p_price" json:"price"`
	IsActive   bool      `gorm:"p_is_active" json:"is_active"`
	CreateDate time.Time `gorm:"column:p_create_date" json:"create_date"`
}
type Count struct {
	Count int `gorm:"column:count" json:"count"`
}
