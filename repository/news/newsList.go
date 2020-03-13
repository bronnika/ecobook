package news

import (
	"../../db"
	"../../models"
)

func NewsList(response *[]models.NewsResponse) error {
	sqlQuery := ``

	if err := db.GetDBConn().Raw(sqlQuery).Scan(&response).Error; err != nil {
		return err
	}

	return nil
}
