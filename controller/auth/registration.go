package auth

import (
	"ecobook/models"
	"ecobook/service/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Registration(c *gin.Context) {
	var (
		request models.Registration
		userId  int
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Registration func bind error:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "something went wrong",
		})
		return
	}

	if err := auth.Registration(request, &userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userId,
	})
}
