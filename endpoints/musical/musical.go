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

	var schedules []models.MusicalSchedule

	for _, schedule := range performance.Schedules {
		schedules = append(schedules, models.MusicalSchedule{
			UUID: schedule.UUID,
			Date: schedule.Date,
			Time: schedule.Time.String,
			Cast: nil,
		})
	}

	musical := models.Musical{
		Id:          performance.ID,
		Title:       performance.Title,
		ThumbUrl:    performance.GetFileURL(),
		RunningTime: performance.RunningTime,
		StartDate:   performance.StartDate,
		EndDate:     performance.EndDate,
		Schedules:   getSchedules(performance.ID),
		Cast:        getCasts(performance.ID),
		Contents:    nil,
	}

	c.JSON(200, musical)
}

func getCasts(id uint) []models.Cast {
	casts := repositories.performance.GetCastsByID(id)
	var result []models.Cast

	for _, cast := range casts {
		result = append(result, models.Cast{
			Id: cast.ID,
			Profile: models.Profile{
				Url: cast.ProfileImageURL(),
			},
			Name: cast.Person.Name,
			Role: cast.Character.Name,
		})
	}

	return result
}

func getSchedules(id uint) []models.MusicalSchedule {
	schedules := repositories.performance.GetSchedulesByID(id)
	if schedules == nil {
		return nil
	}

	var result []models.MusicalSchedule

	for _, schedule := range schedules {
		var casts []models.MusicalScheduleCast

		for _, cast := range schedule.Casts {
			casts = append(casts, models.MusicalScheduleCast{
				ID:   cast.ID,
				Name: cast.Person.Name,
			})
		}

		result = append(result, models.MusicalSchedule{
			UUID: schedule.UUID,
			Date: schedule.Date,
			Time: schedule.Time.String,
			Cast: casts,
		})
	}

	return result
}
