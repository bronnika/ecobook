package news

import (
	"../../db"
	"../../models"
)

func NewsList(response *[]models.NewsResponse) error {
	sqlQuery := `select * from news_list()`

	if err := db.GetDBConn().Raw(sqlQuery).Scan(&response).Error; err != nil {
		return err
	}

	return nil
}
