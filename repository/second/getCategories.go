package second

import (
	"ecobook/db"
	"ecobook/models"
)

func GetCategories(response *[]models.ProductCategory) error {
	sqlQuery := `select id, "name", photo from product_category`

	if err := db.GetDBConn().Raw(sqlQuery).Scan(&response).Error; err != nil {
		return err
	}

	return nil
}
