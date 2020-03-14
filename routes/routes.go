package routes

import (
	"fmt"
	"io"
	"log"
	"os"

	"../controller"
	"../controller/news"
	"../utils"
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
	r.GET("/image", controller.GetImages)

	_ = r.Run(utils.AppSettings.AppParams.PortRun)
}
