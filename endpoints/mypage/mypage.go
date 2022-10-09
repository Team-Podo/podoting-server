package mypage

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/mypage"
	"github.com/gin-gonic/gin"
	"net/http"
)

var repositories Repository

type Repository struct {
	order models.OrderRepository
}

func init() {
	repositories = Repository{
		order: &repository.OrderRepository{DB: database.Gorm},
	}
}

func GetOrderHistory(c *gin.Context) {
	userUID, exists := c.Get("UUID")
	if !exists {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	orders := repositories.order.GetByUserUID(userUID.(string))
	if orders == nil {
		c.JSON(http.StatusNotFound, "Order Not Found")
	}

	c.JSON(200, gin.H{
		"message": "success",
		"orders":  mypage.ParseOrder(orders),
	})
}
