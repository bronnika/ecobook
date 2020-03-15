package utilize_point

import (
	"ecobook/models"
	"ecobook/repository/utilize_point"
)

func GetPoint(categoryID int, response *[]models.UtilizePoint) error {
	return utilize_point.GetPoint(categoryID, response)
}
