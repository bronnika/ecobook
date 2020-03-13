package routes

import (
	"fmt"
	"io"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"../utils"
)

func RunAllRoutes() {
	r := gin.Default()

	f, err := os.Create("gin.log")

	if err != nil {
		fmt.Println("file create error", err.Error())
	}

	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	_ = r.Run(utils.AppSettings.AppParams.PortRun)
}
