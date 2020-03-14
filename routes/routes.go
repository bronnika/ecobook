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
	"ecobook/utils"

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
	r.GET("/my_product_list", my_product_list.MyProductList)

	_ = r.Run(utils.AppSettings.AppParams.PortRun)
}
