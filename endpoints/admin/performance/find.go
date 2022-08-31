package performance

import (
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Find(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	performance := repositories.performance.FindByID(uint(intId))

	if performance == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, getResponseFormForFind(performance))
}

func getResponseFormForFind(p *repository.Performance) utils.MapSlice {
	var performanceProduct utils.MapSlice

	if p.Product.IsNotNil() {
		performanceProduct = utils.BuildMapSliceByMap(map[string]any{
			"id":    p.Product.ID,
			"title": p.Product.Title,
		})
	}

	var schedules []utils.MapSlice

	for _, s := range p.Schedules {
		schedules = append(schedules, utils.BuildMapSliceByMap(map[string]any{
			"uuid": s.UUID,
			"memo": s.Memo,
			"date": s.Date,
			"time": s.Time,
		}))
	}

	return utils.BuildMapSliceByMap(map[string]any{
		"id":          p.ID,
		"title":       p.Title,
		"runningTime": p.RunningTime,
		"startDate":   p.StartDate,
		"endDate":     p.EndDate,
		"rating":      p.Rating,
		"product":     performanceProduct,
		"schedules":   schedules,
		"createdAt":   p.CreatedAt,
		"updatedAt":   p.UpdatedAt,
	})
}
