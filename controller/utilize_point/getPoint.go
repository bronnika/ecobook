package utilize_point

import (
	"ecobook/models"
	"ecobook/service/utilize_point"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPoint(c *gin.Context) {
	var (
		//categoryID = c.Param("category_id")
		response []models.UtilizePoint
	)

	if err := utilize_point.GetPoint(1, &response); err != nil {
		log.Println("GetUtilizePoint error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	log.Println("GetUtilizePoint", http.StatusOK)
	c.JSON(http.StatusOK, response)
}
