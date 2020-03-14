package auth

import (
	"ecobook/models"
	"ecobook/service/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func OtpServer(c *gin.Context) {
	var request models.OtpServer

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("OtpServer func bind error:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "something went wrong",
		})
		return
	}

	log.Println(request)

	if err := auth.OtpServer(request.PhoneNum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reason": "OTP отправлен на номер " + request.PhoneNum,
	})
}
