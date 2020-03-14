package models

type ProductCategory struct {
	ID    int    `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Photo string `json:"photo" gorm:"column:photo"`
}
