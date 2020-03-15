package my_product_list

import (
	"ecobook/models"
	"ecobook/service/my_product_list"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FavoriteProducts(c *gin.Context) {
	var (
		userID   = c.GetHeader("user_id")
		response []models.FavoriteProduct
	)

	if err := my_product_list.FavoriteProducts(userID, &response); err != nil {
		log.Println("FavoriteProducts error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	log.Println("FavoriteProducts", http.StatusOK, userID)
	c.JSON(http.StatusOK, response)
}
