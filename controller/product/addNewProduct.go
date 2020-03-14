package product

import (
	"ecobook/models"
	"ecobook/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func AddNewProduct(c *gin.Context) {
	var newProduct models.AddNewProductReq
	categoryID := c.PostForm("category_id")
	userID := c.PostForm("user_id")
	newProduct.Name = c.PostForm("name")
	newProduct.Description = c.PostForm("description")
	saleTypeID := c.PostForm("sale_type_id")
	price := c.PostForm("price")

	newProduct.CategoryID, _ = strconv.ParseInt(categoryID, 10, 64)
	newProduct.UserID, _ = strconv.ParseInt(userID, 10, 64)
	newProduct.SaleTypeID, _ = strconv.ParseInt(saleTypeID, 10, 64)
	newProduct.Price, _ = strconv.ParseFloat(price, 64)

	form, _ := c.MultipartForm()

	files := form.File["photo"]
	file := files[0]

	newProduct.Photo = userID + time.Now().Format("_02_01_2006_15_04_05") + ".jpg"
	err := service.AddNewProduct(newProduct)

	if err != nil {
		c.JSON(400, gin.H{"reason": "something want wrong"})
		return
	}

	path := "files/" + newProduct.Photo

	if err = c.SaveUploadedFile(file, path); err != nil {
		log.Println("error on saving file -> ", err.Error())
		c.JSON(400, gin.H{"reason": "error on saving image"})
		return
	}

	c.JSON(200, gin.H{"reason": "success"})
}

func GetSaleTypes(c *gin.Context) {
	response := service.GetSaleTypes()
	c.JSON(200, response)
}
