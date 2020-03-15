package models

type UtilizePoint struct {
	ID          int    `json:"id" gorm:"column:id"`
	Name        string `json:"name" gotm:"column:name"`
	Description string `json:"decription" gorm:"column:description"`
	Address     string `json:"address" gorm:"column:address"`
	Phone       string `json:"phone" gorm:"column:phone"`
	Latitude    string `json:"latitude" gorm:"column:latitude"`
	Longitude   string `json:"longitude" gorm:"column:longitude"`
	Photo       string `json:"photo" gorm:"column:photo"`
}

type UtilizeCategory struct {
	ID            int            `json:"id" gorm:"column:id"`
	Name          string         `json:"name" gorm:"column:name"`
	Photo         string         `json:"photo" gorm:"column:photo"`
	UtilizePoints []UtilizePoint `json:"utilize_points"`
}

type UtilizePointShortDesc struct {
	ID    int    `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Photo string `json:"photo" gorm:"column:photo"`
}
