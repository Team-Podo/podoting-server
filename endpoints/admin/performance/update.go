package performance

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(c *gin.Context) {
	var json request

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	performance := repositories.performance.FindByID(id)
	if performance == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "performance not found",
		})
		return
	}

	fmt.Println(&json.MainAreaID)

	performance.PlaceID = &json.PlaceID
	performance.MainAreaID = func() *uint {
		if json.MainAreaID != 0 {
			return &json.MainAreaID
		}
		return nil
	}()
	performance.Title = json.Title
	performance.RunningTime = json.RunningTime
	performance.StartDate = json.StartDate
	performance.EndDate = json.EndDate
	performance.Rating = json.Rating

	err = repositories.performance.Update(performance)

	if err != nil {
		c.JSON(http.StatusNotFound, "Performance Not Found Please Check ID")
		return
	}

	c.JSON(http.StatusOK, performance.ID)
}
