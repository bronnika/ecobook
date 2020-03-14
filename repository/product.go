package repository

import (
	"ecobook/db"
	"ecobook/models"
	"log"
)

func GetProductsByCategoryID(categoryID int64) []models.Product {
	var products []models.Product
	sqlQuery := `SELECT pr.*, u.name as user_name, u.phone as user_phone, st.name as sale_type_name from product as pr
inner join public.user as u on pr.user_id = u.id
inner join sale_type st on pr.sale_type_id = st.id
where pr.category_id = ?`

	if err := db.GetDBConn().Raw(sqlQuery, categoryID).Scan(&products).Error; err != nil {
		log.Println("GetProductsByCategoryID -> ", err)
		return products
	}
	return products
}
