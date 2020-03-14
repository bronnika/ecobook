package auth

import (
	"ecobook/models"
	"ecobook/service/auth"
	"ecobook/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func OtpCheck(c *gin.Context) {
	var (
		request      models.OtpCheck
		token        = utils.AppSettings.Token
		isRegistered bool
		userId       int
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("OtpCheck func bind error:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "something went wrong",
		})
		return
	}

	if err := auth.OtpCheck(request.PhoneNum, request.OtpCode, &isRegistered, &userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":       userId,
		"token":         token,
		"is_registered": isRegistered,
	})
}
