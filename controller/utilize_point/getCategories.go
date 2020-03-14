package utilize_point

import (
	"ecobook/models"
	"ecobook/service/utilize_point"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var (
		response []models.UtilizeCategory
	)

	if err := utilize_point.GetCategories(&response); err != nil {
		log.Println("GetUtilizeCategories error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	log.Println("GetUtilizeCategories", http.StatusOK)
	c.JSON(http.StatusOK, response)
}
