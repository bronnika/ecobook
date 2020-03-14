package my_product_list

import (
	"ecobook/db"
)

func LikeProduct(userId string, productId string) error {
	sqlQuery := `select * from like_product(?,?)`

	if err := db.GetDBConn().Exec(sqlQuery, userId, productId).Error; err != nil {
		return err
	}
	return nil
}
