package performance

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	err = repositories.performance.Delete(uint(intId))
	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, nil)
}
