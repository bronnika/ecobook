package utilize_point

import (
	"ecobook/db"
	"ecobook/models"
)

func GetPoints(categoryId string, response *[]models.UtilizePointShortDesc) error {
	sqlQuery := `select id, "name", photo from utilize_point where category_id = ?`

	if err := db.GetDBConn().Raw(sqlQuery, categoryId).Scan(&response).Error; err != nil {
		return err
	}
	return nil
}
