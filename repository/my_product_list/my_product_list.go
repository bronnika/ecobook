package my_product_list

import (
	"ecobook/db"
	"ecobook/models"
)

func MyProductList(response *[]models.MyProductList, userID int) error {
	sqlQuery := `select * from my_product_list(?)`

	if err := db.GetDBConn().Raw(sqlQuery, userID).Scan(&response).Error; err != nil {
		return err
	}

	return nil
}
