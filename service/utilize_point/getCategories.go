package utilize_point

import (
	"ecobook/models"
	"ecobook/repository/utilize_point"
	"log"
)

func GetCategories(response *[]models.UtilizeCategory) error {
	if err := utilize_point.GetCategories(response); err != nil {
		return err
	}

	for i, row := range *response {
		var untilizePoints []models.UtilizePoint

		if err := GetPoint(row.ID, &untilizePoints); err != nil {
			log.Println("GetCategories func get point error:", err.Error)
		}

		row.UtilizePoints = untilizePoints

		(*response)[i] = row
	}

	log.Println("response:", response)

	return nil
}
