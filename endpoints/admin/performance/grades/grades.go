package grades

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/admin/seat_grade_get"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var repositories Repository

type request struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Price int    `json:"price"`
}

type Repository struct {
	seatGrade models.SeatGradeRepository
}

func init() {
	repositories = Repository{
		seatGrade: &repository.SeatGradeRepository{DB: database.Gorm},
	}
}

func Get(c *gin.Context) {
	seatGrades, err := repositories.seatGrade.GetAll()
	if len(seatGrades) == 0 || err != nil {
		c.JSON(http.StatusNotFound, "SeatGrade not found")
		return
	}

	c.JSON(http.StatusOK, seatGrades)
}

func GetByPerformanceID(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	seatGrades, err := repositories.seatGrade.GetByPerformanceID(performanceID)
	if len(seatGrades) == 0 || err != nil {
		c.JSON(http.StatusNotFound, "SeatGrade not found")
		return
	}

	c.JSON(http.StatusOK, seat_grade_get.ParseResponseForm(seatGrades))
}

func Create(c *gin.Context) {
	PerformanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	var req request
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	SeatGrade := &repository.SeatGrade{}
	SeatGrade.Name = req.Name
	SeatGrade.Color = req.Color
	SeatGrade.Price = req.Price
	SeatGrade.PerformanceID = PerformanceID

	err = repositories.seatGrade.Create(SeatGrade)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, SeatGrade.ID)
}

func Update(c *gin.Context) {
	SeatGradeID, err := utils.ParseUint(c.Param("seat_grade_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(SeatGrade) id should be Integer")
		return
	}

	var req request
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	SeatGrade := &repository.SeatGrade{}
	SeatGrade.ID = SeatGradeID
	SeatGrade.Name = req.Name
	SeatGrade.Color = req.Color
	SeatGrade.Price = req.Price

	err = repositories.seatGrade.Update(SeatGrade)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, SeatGrade.ID)
}

func Delete(c *gin.Context) {
	SeatGradeID, err := utils.ParseUint(c.Param("seat_grade_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(SeatGrade) id should be Integer")
		return
	}

	err = repositories.seatGrade.Delete(&repository.SeatGrade{
		ID: SeatGradeID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, SeatGradeID)
}
