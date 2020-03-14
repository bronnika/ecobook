package utilize_point

import (
	"ecobook/models"
	"ecobook/repository/utilize_point"
)

func GetCategories(response *[]models.UtilizeCategory) error {
	return utilize_point.GetCategories(response)
}
