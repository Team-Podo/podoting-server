package performance

import (
	"github.com/Team-Podo/podoting-server/admin/performance/cast"
	"github.com/gin-gonic/gin"
)

func SetRoute(admin *gin.RouterGroup) {
	performances := admin.Group("/performances")

	cast.SetRoute(performances)
}
