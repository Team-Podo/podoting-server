package order

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var repositories Repository

type Repository struct {
	order       models.OrderRepository
	orderDetail models.OrderDetailRepository
}

func init() {
	repositories = Repository{
		order:       &repository.OrderRepository{DB: database.Gorm},
		orderDetail: &repository.OrderDetailRepository{DB: database.Gorm},
	}
}

func CancelOrder(c *gin.Context) {
	userUID, exists := c.Get("UUID")
	if !exists {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	orderID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(order) id must be a integer")
		return
	}

	order := repositories.order.FindByID(orderID)
	if order == nil {
		c.JSON(http.StatusNotFound, "Order Not Found")
		return
	}

	if order.BuyerUID != userUID.(string) {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	err = repositories.order.CancelOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func CancelOrderDetail(c *gin.Context) {
	userUID, exists := c.Get("UUID")
	if !exists {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	orderID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(order) id must be a integer")
		return
	}

	orderDetailID, err := utils.ParseUint(c.Param("order_detail_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(order_detail) id must be a integer")
		return
	}

	order := repositories.order.FindByID(orderID)
	if order == nil {
		c.JSON(http.StatusNotFound, "Order Not Found")
		return
	}

	if order.BuyerUID != userUID.(string) {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	orderDetail := repositories.orderDetail.FindByID(orderDetailID)
	if orderDetail == nil {
		c.JSON(http.StatusNotFound, "Order Detail Not Found")
		return
	}

	err = repositories.orderDetail.CancelOrderDetail(orderDetail)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
