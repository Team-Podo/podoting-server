package book

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var repositories Repository

type Repository struct {
	seatBooking models.SeatBookingRepository
	schedule    models.ScheduleRepository
	order       models.OrderRepository
}

func init() {
	repositories = Repository{
		seatBooking: &repository.SeatBookingRepository{DB: database.Gorm},
		schedule:    &repository.ScheduleRepository{DB: database.Gorm},
		order:       &repository.OrderRepository{DB: database.Gorm},
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

	schedule := repositories.schedule.Find(scheduleUUID)
	if schedule == nil {
		c.JSON(http.StatusNotFound, "Schedule Not Found")
		return
	}

	userUID, exists := c.Get("UUID")
	if !exists {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	err := repositories.seatBooking.Book(userUID.(string), scheduleUUID, requests.SeatUUIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	seatBookings, err := repositories.seatBooking.Get(userUID.(string), scheduleUUID, requests.SeatUUIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var details []repository.OrderDetail
	for _, seatBooking := range seatBookings {
		details = append(details, repository.OrderDetail{
			SeatBookingID:  seatBooking.ID,
			OriginalPrice:  uint(seatBooking.Seat.Grade.Price),
			OrderDetailKey: utils.GenerateOrderDetailKey(),
		})
	}

	var order repository.Order
	order.PerformanceID = schedule.PerformanceID
	order.ScheduleUUID = scheduleUUID
	order.OrderKey = utils.GenerateOrderKey()
	order.BuyerUID = userUID.(string)
	order.Paid = true
	order.Details = details

	err = repositories.order.Save(&order)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
