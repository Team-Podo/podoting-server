package performance

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/repository"
	response "github.com/Team-Podo/podoting-server/response/admin/performance_get"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Get(c *gin.Context) {
	queryMap := extractQueriesToMap(c)

	performances := repositories.performance.GetWithQueryMap(queryMap)

	if performances == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	fmt.Println(performances)

	c.JSON(http.StatusOK, gin.H{
		"performances": getResponseFormForGet(performances),
		"total":        repositories.performance.GetTotalWithQueryMap(queryMap),
	})
}

func getResponseFormForGet(ps []repository.Performance) []response.Performance {
	var res []response.Performance

	for _, p := range ps {
		res = append(res, response.Performance{
			ID:          p.ID,
			Title:       p.Title,
			RunningTime: p.RunningTime,
			StartDate:   p.StartDate,
			EndDate:     p.EndDate,
			Rating:      p.Rating,
			Product:     getProductFromPerformance(&p),
			Schedules:   p.Schedules,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	return res
}

func getProductFromPerformance(p *repository.Performance) *response.Product {
	if p.Product == nil {
		return nil
	}

	product := p.Product

	return &response.Product{
		ID:        product.ID,
		Title:     product.Title,
		File:      getFileFromProduct(product),
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func getFileFromProduct(p *repository.Product) *response.File {
	if p.File == nil {
		return nil
	}

	file := p.File

	return &response.File{
		ID:        file.ID,
		Size:      file.Size,
		Path:      file.Path,
		CreatedAt: file.CreatedAt,
		UpdatedAt: file.UpdatedAt,
	}
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
