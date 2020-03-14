package my_product_list

import "ecobook/repository/my_product_list"

func LikeProduct(userId string, productId string) error {
	return my_product_list.LikeProduct(userId, productId)
}
