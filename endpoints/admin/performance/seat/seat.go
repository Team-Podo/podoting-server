package seat

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/admin/seat_get"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	UUID    string `json:"uuid" binding:"required"`
	GradeID uint   `json:"gradeID" binding:"required"`
}

type Repository struct {
	seat models.SeatRepository
}

func init() {
	repositories = Repository{
		seat: &repository.SeatRepository{DB: database.Gorm},
	}
}

var repositories Repository

func Get(c *gin.Context) {
	performanceID, err := getPerformanceID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	areaID, err := getAreaID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "(area) id should be Integer")
		return
	}

	seats := repositories.seat.GetByAreaAndPerformanceID(*areaID, *performanceID)

	c.JSON(http.StatusOK, seat_get.ParseResponseForm(seats))
}

func Save(c *gin.Context) {
	performanceID, err := getPerformanceID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	var requests []Request
	if err = c.ShouldBindJSON(&requests); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	var seats []repository.Seat
	for _, request := range requests {
		seats = append(seats, repository.Seat{
			PerformanceID: *performanceID,
			UUID:          request.UUID,
			SeatGradeID:   request.GradeID,
		})
	}

	if err = repositories.seat.SaveSeats(seats); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "success")
}

func getPerformanceID(c *gin.Context) (*uint, error) {
	id, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func getAreaID(c *gin.Context) (*uint, error) {
	id, err := utils.ParseUint(c.Param("area_id"))
	if err != nil {
		return nil, err
	}

	return &id, nil
}
