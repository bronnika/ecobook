package second

import (
	"log"
	"net/http"

	"ecobook/models"
	"ecobook/service/second"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var (
		response []models.ProductCategory
	)

	if err := second.GetCategories(&response); err != nil {
		log.Println("GetProductCategories error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	log.Println("GetProductCategory", http.StatusOK)
	c.JSON(http.StatusOK, response)
}
