package my_product_list

import (
	"ecobook/models"
	"ecobook/repository/my_product_list"
)

func MyProductList(response *[]models.MyProductList, userID int) error {
	return my_product_list.MyProductList(response, userID)
}
