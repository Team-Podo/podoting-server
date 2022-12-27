package cast

import (
	"github.com/Team-Podo/podoting-server/admin/performance/cast/request"
	"github.com/Team-Podo/podoting-server/admin/performance/cast/response"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var repositories Repository

type Repository struct {
	cast        models.CastRepository
	performance models.PerformanceRepository
}

func init() {
	repositories = Repository{
		cast:        &repository.CastRepository{DB: database.Gorm},
		performance: &repository.PerformanceRepository{DB: database.Gorm},
	}
}

func Index(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	performance := repositories.performance.FindByID(performanceID)
	if performance == nil {
		c.JSON(http.StatusNotFound, "performance not found")
		return
	}

	casts, err := repositories.cast.FindByPerformanceID(performanceID)
	if err != nil {
		c.JSON(http.StatusNotFound, "casts not found")
		return
	}

	c.JSON(http.StatusOK, response.GetIndexResponse(casts, performance))
}

func SaveManyCasts(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	var createManyCastRequests []request.CreateManyCastRequest
	err = c.BindJSON(&createManyCastRequests)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON")
		return
	}

	var casts []repository.Cast

	for _, createManyCastRequest := range createManyCastRequests {
		var newCast repository.Cast

		if createManyCastRequest.ID != 0 {
			cast, err := repositories.cast.FindOneByID(createManyCastRequest.ID)
			if err != nil {
				c.JSON(http.StatusNotFound, "cast not found")
				return
			}

			newCast = *cast
		}

		newCast.PerformanceID = performanceID
		newCast.PersonID = createManyCastRequest.PersonID
		newCast.CharacterID = createManyCastRequest.CharacterID

		casts = append(casts, newCast)
	}

	err = repositories.cast.CreateMany(casts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "database error: cast create failed")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"casts": response.GetCreateManyResponse(casts),
	})
}
