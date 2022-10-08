package seat

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/musical/seat"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var repositories Repository

type Repository struct {
	seat        models.SeatRepository
	performance models.PerformanceRepository
	area        models.AreaRepository
	schedule    models.ScheduleRepository
}

func init() {
	repositories = Repository{
		seat:        &repository.SeatRepository{DB: database.Gorm},
		performance: &repository.PerformanceRepository{DB: database.Gorm},
		area:        &repository.AreaRepository{DB: database.Gorm},
		schedule:    &repository.ScheduleRepository{DB: database.Gorm},
	}
}

func Get(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	areaID := getAreaIDByUint(c)

	if areaID == 0 {
		mainAreaId, err := repositories.performance.CheckMainAreaExistsByID(parseParamToUint(c, "id"))
		if err != nil {
			c.JSON(http.StatusNotFound, "Main Area Not Found")
			return
		}

		areaID = mainAreaId
	}

	scheduleUUID := c.Param("schedule_uuid")

	backgroundImage := repositories.area.GetBackgroundImageByAreaId(areaID)
	seats := repositories.seat.GetSeatsByAreaIdAndScheduleUUID(areaID, scheduleUUID)
	schedules, err := repositories.schedule.FindByPerformanceID(performanceID)
	if err != nil {
		c.JSON(http.StatusNotFound, "Schedule Not Found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"backgroundImage": backgroundImage,
		"seats":           seat.ParseSeats(seats),
		"schedules":       seat.ParseSchedules(schedules),
	})
}

func getAreaIDByUint(c *gin.Context) uint {
	areaID, err := strconv.Atoi(c.Query("area_id"))
	if err != nil {
		return 0
	}

	return uint(areaID)
}

func parseParamToUint(c *gin.Context, param string) uint {
	id, err := strconv.Atoi(c.Param(param))

	if err != nil {
		return 0
	}

	return uint(id)
}
