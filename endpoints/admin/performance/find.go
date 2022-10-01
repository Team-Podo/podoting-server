package performance

import (
	response "github.com/Team-Podo/podoting-server/response/admin/performance_find"
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
		c.JSON(http.StatusNotFound, "performance not found please check id")
		return
	}

	c.JSON(http.StatusOK, response.ParseResponseForm(performance))
}
