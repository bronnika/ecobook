package participate

import (
	"ecobook/service/participate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Participate(c *gin.Context) {
	eventID, err1 := strconv.Atoi(c.Query("event_id"))
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil || err1 != nil {
		log.Println("Participate error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}
	count, err := participate.Participate(eventID, userID)
	if err != nil {
		log.Println("Participate error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "что-то пошло не так",
		})
		return
	}

	log.Println(http.StatusOK, "Participate")
	c.JSON(http.StatusOK, gin.H{"count": count, "reason": "Успешно присоединились"})
}
