package models

type UtilizePoint struct {
}

type UtilizeCategory struct {
	ID    int    `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Photo string `json:"photo" gorm:"column:photo"`
}

type UtilizePointShortDesc struct {
	ID    int    `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Photo string `json:"photo" gorm:"column:photo"`
}
