package book

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

var repositories Repository

type Repository struct {
	seatBooking models.SeatBookingRepository
}

func init() {
	repositories = Repository{
		seatBooking: &repository.SeatBookingRepository{DB: database.Gorm},
	}
}

type Request struct {
	SeatUUIDs []string `json:"seat_uuids"`
}

func Book(c *gin.Context) {
	scheduleUUID := c.Param("schedule_uuid")

	var requests Request
	if err := c.ShouldBindJSON(&requests); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	err := repositories.seatBooking.Book(scheduleUUID, requests.SeatUUIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
