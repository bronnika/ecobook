package my_product_list

import (
	"ecobook/db"
	"ecobook/models"
)

func FavoriteProducts(userId string, response *[]models.FavoriteProduct) error {
	sqlQuery := `select * from favorite_product_list(?)`

	if err := db.GetDBConn().Raw(sqlQuery, userId).Scan(&response).Error; err != nil {
		return err
	}

	return nil
}
