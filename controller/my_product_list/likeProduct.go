package my_product_list

import (
	"ecobook/service/my_product_list"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LikeProduct(c *gin.Context) {
	var (
		userID    = c.GetHeader("user_id")
		productID = c.Param("product_id")
	)

	if err := my_product_list.LikeProduct(userID, productID); err != nil {
		log.Println("LikeProduct error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	log.Println(http.StatusOK, userID, productID)
	c.JSON(http.StatusOK, gin.H{
		"reason": "success",
	})
}
