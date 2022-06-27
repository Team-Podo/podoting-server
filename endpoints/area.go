package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/database"
	"github.com/kwanok/podonine/repository"
	"net/http"
	"strconv"
)

func GetArea(c *gin.Context) {
	id := c.Param("areaId")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		c.Abort()
		return
	}

	areaRepository := repository.AreaRepository{Db: database.Gorm}
	area := areaRepository.Find(uint(intId))

	if area == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, area)
}
