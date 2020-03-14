package routes

import (
	"fmt"
	"io"
	"log"
	"os"

	"ecobook/controller"
	"ecobook/controller/news"
	"ecobook/controller/second"
	"ecobook/controller/utilize_point"
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

	utilize := r.Group("/utilize")
	utilize.GET("/categories", utilize_point.GetCategories)
	utilize.GET("/categories/:category_id", utilize_point.GetPoints)
	utilize.GET("/point/:id", utilize_point.GetPoint)

	_ = r.Run(utils.AppSettings.AppParams.PortRun)
}
