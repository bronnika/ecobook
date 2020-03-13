package news

import (
	"../../models"
	"../../repository/news"
)

func NewsList(response *[]models.NewsResponse) error {
	return news.NewsList(response)
}
