package performance

import (
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
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

	c.JSON(http.StatusOK, gin.H{
		"performances": getResponseFormForGet(performances),
		"total":        repositories.performance.GetTotalWithQueryMap(queryMap),
	})
}

func getResponseFormForGet(ps []repository.Performance) []utils.MapSlice {
	var res []utils.MapSlice

	for _, p := range ps {
		res = append(res, utils.BuildMapSliceByMap(map[string]any{
			"id":          p.ID,
			"title":       p.Title,
			"runningTime": p.RunningTime,
			"startDate":   p.StartDate,
			"endDate":     p.EndDate,
			"rating":      p.Rating,
			"product":     p.Product,
			"schedules":   p.Schedules,
			"createdAt":   p.CreatedAt,
			"updatedAt":   p.UpdatedAt,
		}))
	}

	return res
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
