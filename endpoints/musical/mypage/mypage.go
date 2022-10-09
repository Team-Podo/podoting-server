package mypage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrderHistory(c *gin.Context) {
	userUID, exists := c.Get("UUID")
	if !exists {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	fmt.Println(userUID)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
