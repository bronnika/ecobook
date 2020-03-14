package routes

import (
	"ecobook/controller/my_product_list"
	"fmt"
	"io"
	"log"
	"os"

	"ecobook/controller"
	"ecobook/controller/news"
	"ecobook/controller/second"
	"ecobook/controller/utilize_point"
	"ecobook/utils"

	"ecobook/controller/product"
	"github.com/gin-gonic/gin"
)

func RunAllRoutes() {
	r := gin.Default()

	f, err := os.Create("gin.log")

	if err != nil {
		fmt.Println("file create error", err.Error())
	}

	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.GET("/news_list", news.NewsList)
	r.GET("/image/:image_name", controller.GetImages)
	r.GET("/product_categories", second.GetCategories)
	r.GET("/products/:categoryID", product.GetProductList)
	r.GET("/my_product_list", my_product_list.MyProductList)

	utilize := r.Group("/utilize")
	utilize.GET("/categories", utilize_point.GetCategories)
	utilize.GET("/categories/:category_id", utilize_point.GetPoints)
	utilize.GET("/point/:id", utilize_point.GetPoint)

	r.GET("/my_product_list", my_product_list.MyProductList)

	_ = r.Run(utils.AppSettings.AppParams.PortRun)
}
