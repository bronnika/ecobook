package utilize_point

import (
	"ecobook/db"
	"ecobook/models"
)

func GetPoint(pointId string, response *models.UtilizePoint) error {
	sqlQuery := `select * from utilize_point where id = ?`

	if err := db.GetDBConn().Raw(sqlQuery, pointId).Scan(&response).Error; err != nil {
		return err
	}
	return nil
}
