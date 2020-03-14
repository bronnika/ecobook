package routes

import (
	"ecobook/controller/auth"
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

	r.POST("/otp", auth.OtpServer)
	r.POST("/otp_check", auth.OtpCheck)

	r.POST("/registration", auth.Registration)

	r.GET("ws")

	r.GET("/news_list", news.NewsList)
	r.GET("/image/:image_name", controller.GetImages)
	r.GET("/product_categories", second.GetCategories)

	go controller.HandleMessages()

	_ = r.Run(utils.AppSettings.AppParams.PortRun)
}
