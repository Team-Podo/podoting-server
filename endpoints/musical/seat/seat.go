package seat

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/musical/seat"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var repositories Repository

type Repository struct {
	seat        models.SeatRepository
	performance models.PerformanceRepository
	area        models.AreaRepository
}

func init() {
	repositories = Repository{
		seat:        &repository.SeatRepository{DB: database.Gorm},
		performance: &repository.PerformanceRepository{DB: database.Gorm},
		area:        &repository.AreaRepository{DB: database.Gorm},
	}
}

func Get(c *gin.Context) {
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

	seats := repositories.seat.GetSeatsByAreaIdAndScheduleUUID(areaID, scheduleUUID)
	backgroundImage := repositories.area.GetBackgroundImageByAreaId(areaID)

	response := make([]seat.Seat, len(seats))

	for i := range seats {
		Booked := false
		s := seats[i]

		if len(seats[i].Bookings) > 0 {
			Booked = true
		}

		response[i] = seat.Seat{
			UUID:   s.UUID,
			Grade:  seat.Grade{ID: s.Grade.ID, Name: s.Grade.Name},
			Price:  s.Grade.Price,
			Color:  s.Grade.Color,
			Booked: Booked,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"backgroundImage": backgroundImage,
		"seats":           response,
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
