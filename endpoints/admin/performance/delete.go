package performance

import (
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Delete(c *gin.Context) {
	id, err := utils.ParseUint(c.Param("id"))

	err = repositories.performance.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
