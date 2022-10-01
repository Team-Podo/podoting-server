package performance

import (
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
		c.JSON(http.StatusNotFound, "performances not found")
		return
	}

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
			ThumbUrl:    getThumbUrl(&p),
			RunningTime: p.RunningTime,
			StartDate:   p.StartDate,
			EndDate:     p.EndDate,
			Rating:      p.Rating,
			Schedules:   p.Schedules,
			CreatedAt:   p.CreatedAt.String(),
			UpdatedAt:   p.UpdatedAt.String(),
		})
	}

	return res
}

func getThumbUrl(p *repository.Performance) *string {
	if p.Thumbnail != nil {
		thumbUrl := p.Thumbnail.FullPath()
		return &thumbUrl
	}

	return nil
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
