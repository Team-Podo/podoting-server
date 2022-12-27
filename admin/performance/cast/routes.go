package cast

import "github.com/gin-gonic/gin"

func SetRoute(performances *gin.RouterGroup) {
	casts := performances.Group("/:id/casts")

	casts.GET("/", Index)
	casts.POST("/many", SaveManyCasts)
}
