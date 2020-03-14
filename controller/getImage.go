package controller

import (
	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {
	imageName := c.Param("image_name")

	// returning requested image
	c.File("files/" + imageName)
}