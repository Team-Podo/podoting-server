package mainpage

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/endpoints/mainpage/res"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/gin-gonic/gin"
)

var repositories Repository

type Repository struct {
	performance models.PerformanceRepository
}

func init() {
	repositories = Repository{
		performance: &repository.PerformanceRepository{DB: database.Gorm},
	}
}

func Index(c *gin.Context) {
	performances := repositories.performance.GetWithQueryMap(nil)

	c.JSON(200, res.MainPageResponse{}.Of(performances))
}
