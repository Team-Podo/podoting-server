package performance

import (
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Update(c *gin.Context) {
	var json request

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	performance := repository.Performance{
		ID:          uint(intId),
		PlaceID:     &json.PlaceID,
		Title:       json.Title,
		RunningTime: json.RunningTime,
		StartDate:   json.StartDate,
		EndDate:     json.EndDate,
		Rating:      json.Rating,
	}

	err = repositories.performance.Update(&performance)

	if err != nil {
		c.JSON(http.StatusNotFound, "Performance Not Found Please Check ID")
		return
	}

	c.JSON(http.StatusOK, performance.ID)
}
