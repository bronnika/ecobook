package news

import (
	"ecobook/db"
	"ecobook/models"
)

func NewsList(userId string, response *[]models.NewsResponse) error {
	sqlQuery := `select * from news_list_ext(?)`

	if err := db.GetDBConn().Raw(sqlQuery, userId).Scan(&response).Error; err != nil {
		return err
	}

	return nil
}
