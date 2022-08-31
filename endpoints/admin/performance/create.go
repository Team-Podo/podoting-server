package performance

import (
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var json request
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	performance := repository.Performance{
		Product:     &repository.Product{ID: json.ProductID},
		Place:       &repository.Place{ID: json.PlaceID},
		Title:       json.Title,
		RunningTime: json.RunningTime,
		StartDate:   json.StartDate,
		EndDate:     json.EndDate,
		Rating:      json.Rating,
	}

	err := repositories.performance.Save(&performance)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, performance.ID)
}
