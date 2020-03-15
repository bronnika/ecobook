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

type FavoriteProduct struct {
	ID           int     `json:"id" gorm:"column:p_id"`
	Name         string  `json:"name" gorm:"column:p_name"`
	Photo        string  `json:"photo" gorm:"column:p_photo"`
	Description  string  `json:"description" gorm:"column:p_description"`
	BuyerID      int     `json:"buyer_id" gorm:"column:p_buyer_id"`
	SaleTypeName string  `json:"sale_type_name" gorm:"column:p_sale_type_name"`
	Price        float64 `json:"price" gorm:"column:p_price"`
}

type Count struct {
	Count int `gorm:"column:count" json:"count"`
}
