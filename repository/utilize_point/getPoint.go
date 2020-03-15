package utilize_point

import (
	"ecobook/db"
	"ecobook/models"
)

func GetPoint(categoryID int, response *[]models.UtilizePoint) error {
	sqlQuery := `select * from utilize_point where category_id = ?`

	if err := db.GetDBConn().Raw(sqlQuery, categoryID).Scan(&response).Error; err != nil {
		return err
	}
	return nil
}
