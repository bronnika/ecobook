package my_product_list

import (
	"ecobook/models"
	"ecobook/service/my_product_list"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func MyProductList(c *gin.Context) {
	var (
		response []models.MyProductList
	)

	if err := my_product_list.MyProductList(&response); err != nil {
		log.Println("MyProductList error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	log.Println(http.StatusOK, "MyProductList")
	c.JSON(http.StatusOK, response)
}
