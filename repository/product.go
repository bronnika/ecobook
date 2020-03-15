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
where pr.category_id = ? and pr.is_active = true order by pr.id desc`

	if err := db.GetDBConn().Raw(sqlQuery, categoryID).Scan(&products).Error; err != nil {
		log.Println("GetProductsByCategoryID -> ", err)
		return products
	}
	return products
}

func AddNewProduct(newProduct models.AddNewProductReq) error {
	sqlQuery := `insert into product (category_id, name, photo, description, user_id, 
				sale_type_id, price) values (?, ?, ?, ?, ?, ?, ?)`

	err := db.GetDBConn().Exec(sqlQuery, newProduct.CategoryID, newProduct.Name, newProduct.Photo,
		newProduct.Description, newProduct.UserID, newProduct.SaleTypeID, newProduct.Price).Error
	if err != nil {
		log.Println("error AddNewProduct insert query -> ", err)
		return err
	}
	return nil
}

func GetSalesTypes() []models.SalesTypes {
	sqlQuery := `select * from sale_type order by id asc`
	var saleTypes []models.SalesTypes
	if err := db.GetDBConn().Raw(sqlQuery).Scan(&saleTypes).Error; err != nil {
		log.Println("GetProductsByCategoryID -> ", err)
		return saleTypes
	}
	return saleTypes
}
