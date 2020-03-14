package utilize_point

import (
	"ecobook/db"
	"ecobook/models"
)

func GetCategories(response *[]models.UtilizeCategory) error {
	sqlQuery := `select id, "name", photo from point_category`

	if err := db.GetDBConn().Raw(sqlQuery).Scan(&response).Error; err != nil {
		return err
	}

	return nil
}
