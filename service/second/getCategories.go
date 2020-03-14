package second

import (
	"ecobook/models"
	"ecobook/repository/second"
)

func GetCategories(response *[]models.ProductCategory) error {
	return second.GetCategories(response)
}
