package utilize_point

import (
	"ecobook/models"
	"ecobook/service/utilize_point"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPoints(c *gin.Context) {
	var (
		categoryID = c.Param("category_id")
		response   []models.UtilizePointShortDesc
	)

	if err := utilize_point.GetPoints(categoryID, &response); err != nil {
		log.Println("GetUtilizePoints error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	log.Println("GetUtilizePoints", http.StatusOK)
	c.JSON(http.StatusOK, response)
}
