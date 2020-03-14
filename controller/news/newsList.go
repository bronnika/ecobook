package news

import (
	"log"
	"net/http"

	"ecobook/models"
	"ecobook/service/news"

	"github.com/gin-gonic/gin"
)

func NewsList(c *gin.Context) {
	var (
		response []models.NewsResponse
	)

	if err := news.NewsList(&response); err != nil {
		log.Println("NewsList error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	log.Println(http.StatusOK, "NewsList")
	c.JSON(http.StatusOK, response)
}
