package mypage

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/mypage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	query := extractQueriesToMap(c)

	orders, total := repositories.order.GetByUserUIDWithQuery(userUID.(string), query)
	if orders == nil {
		c.JSON(http.StatusNotFound, "Order Not Found")
	}

	c.JSON(200, gin.H{
		"message": "success",
		"orders":  mypage.ParseOrder(orders),
		"total":   total,
	})
}

func extractQueriesToMap(c *gin.Context) map[string]any {
	query := map[string]any{
		"limit":    getLimitQuery(c),
		"offset":   getOffsetQuery(c),
		"reversed": getReverseQuery(c),
	}

	return query
}

func getLimitQuery(c *gin.Context) int {
	limitQuery := c.Query("limit")
	if limitQuery == "" {
		return 10
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		return 10
	}

	return limit
}

func getOffsetQuery(c *gin.Context) int {
	offsetQuery := c.Query("offset")
	if offsetQuery == "" {
		return 0
	}

	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		return 0
	}

	return offset
}

func getReverseQuery(c *gin.Context) bool {
	reversedQuery := c.Query("reversed")
	if reversedQuery == "" {
		return false
	}

	reversed, err := strconv.ParseBool(reversedQuery)
	if err != nil {
		return false
	}

	return reversed
}
