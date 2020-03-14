package my_product_list

import (
	"ecobook/models"
	"ecobook/repository/my_product_list"
)

func FavoriteProducts(userId string, response *[]models.FavoriteProduct) error {
	return my_product_list.FavoriteProducts(userId, response)
}
