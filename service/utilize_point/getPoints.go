package utilize_point

import (
	"ecobook/models"
	"ecobook/repository/utilize_point"
)

func GetPoints(categoryId string, response *[]models.UtilizePointShortDesc) error {
	return utilize_point.GetPoints(categoryId, response)
}
