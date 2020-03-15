package my_product_list

import (
	"ecobook/models"
	"ecobook/service/my_product_list"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MyProductList(c *gin.Context) {
	var (
		response []models.MyProductList
	)
	userID, _ := strconv.Atoi(c.GetHeader("user_id"))

	if err := my_product_list.MyProductList(&response, userID); err != nil {
		log.Println("MyProductList error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	log.Println(http.StatusOK, "MyProductList")
	c.JSON(http.StatusOK, response)
}
