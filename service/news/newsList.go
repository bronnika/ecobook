package news

import (
	"ecobook/models"
	"ecobook/repository/news"
)

func NewsList(userId string, response *[]models.NewsResponse) error {
	return news.NewsList(userId, response)
}
