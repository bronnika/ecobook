package news

import (
	"ecobook/models"
	"ecobook/repository/news"
)

func NewsList(response *[]models.NewsResponse) error {
	return news.NewsList(response)
}
