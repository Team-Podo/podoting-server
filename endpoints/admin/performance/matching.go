package performance

import (
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MatchArea(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	areaID, err := utils.ParseUint(c.Param("area_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(area) id should be Integer")
		return
	}

	performance := repositories.performance.FindByID(performanceID)
	if performance == nil {
		c.JSON(http.StatusNotFound, "performance not found please check id")
		return
	}

	_, err = repositories.performance.CheckMainAreaExistsByID(performanceID)
	if err != nil {
		performance.MainAreaID = &areaID
	}

	ab, err := repositories.areaBoilerplate.GetAreaBoilerplateByAreaID(*performance.MainAreaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "mainAreaID is not valid")
		return
	}

	err = repositories.areaBoilerplate.SaveSeats(ab, performance.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database_error: seats save failed")
		return
	}

	err = repositories.performance.Update(performance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database_error: performance save failed")
		return
	}

	c.JSON(http.StatusOK, "success")
}
