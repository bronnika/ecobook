package product

import (
	"ecobook/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetProductList(c *gin.Context) {
	categoryID := c.Param("categoryID")

	log.Print("categoryID = ", categoryID)

	catID, err := strconv.ParseInt(categoryID, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"reason": "Incorrect category ID"})
		return
	}

	responseObj := service.GetProductsByCategoryID(catID)

	c.JSON(200, responseObj)

}
