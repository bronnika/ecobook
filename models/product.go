package models

type Product struct {
	ID           int64   `json:"id" gorm:"column:id"`
	CategoryID   int64   `json:"category_id" gorm:"column:category_id"`
	Name         string  `json:"name" gorm:"column:name"`
	Photo        string  `json:"photo" gorm:"column:photo"`
	Description  string  `json:"description" gorm:"column:description"`
	UserID       string  `json:"-" gorm:"column:user_id"`
	SaleTypeID   int     `json:"-" gorm:"column:sale_type_id"`
	Price        float64 `json:"price,omitempty" gorm:"column:price"`
	IsActive     bool    `json:"is_active" gorm:"column:is_active"`
	CreateDate   string  `json:"create_date" gorm:"column:create_date"`
	UserPhone    string  `json:"user_phone" gorm:"column:user_phone"`
	UserName     string  `json:"user_name" gorm:"column:user_name"`
	SaleTypeName string  `json:"-" gorm:"column:sale_type_name"`
}

type AddNewProductReq struct {
	Name        string
	Description string
	CategoryID  int64
	UserID      int64
	SaleTypeID  int64
	Price       float64
	Photo       string
}

type SalesTypes struct {
	ID   int64  `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Code string `json:"code" gorm:"column:code"`
}
