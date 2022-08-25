package musical

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func Find(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	performance := repositories.performance.FindByID(uint(intId))
	if performance == nil {
		c.JSON(http.StatusNotFound, "Musical Not Found")
		return
	}

	musical := models.Musical{
		Id:          performance.ID,
		Title:       performance.Title,
		ThumbUrl:    performance.GetFileURL(),
		RunningTime: "",
		StartDate:   performance.StartDate,
		EndDate:     performance.EndDate,
		Schedules:   nil,
		Cast:        nil,
		Contents:    nil,
	}

	schedules := performance.Schedules
	for _, schedule := range schedules {
		musical.Schedules = append(musical.Schedules, models.MusicalSchedule{
			UUID: schedule.UUID,
			Date: schedule.Date,
			Time: schedule.Time.String,
			Cast: nil,
		})
	}

	c.JSON(200, musical)
}
