package schedule

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/admin/schedule_get"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type request struct {
	Memo  string `json:"memo"`
	Open  bool   `json:"open"`
	Date  string `json:"date" binding:"required"`
	Time  string `json:"time"`
	Casts []uint `json:"casts"`
}

type Repository struct {
	schedule    models.ScheduleRepository
	performance models.PerformanceRepository
	cast        models.CastRepository
}

var repositories Repository

func init() {
	repositories = Repository{
		schedule:    &repository.ScheduleRepository{DB: database.Gorm},
		performance: &repository.PerformanceRepository{DB: database.Gorm},
		cast:        &repository.CastRepository{DB: database.Gorm},
	}
}

func Get(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	performance := repositories.performance.FindByID(performanceID)

	schedules, err := repositories.schedule.FindByPerformanceID(performanceID)
	if err != nil {
		c.JSON(http.StatusNotFound, "schedule not found")
		return
	}

	casts, err := repositories.cast.GetByPerformanceID(performanceID)
	if err != nil {
		c.JSON(http.StatusNotFound, "cast not found")
		return
	}

	scheduleRes, castRes := schedule_get.ParseResponseFrom(schedules, casts)

	c.JSON(http.StatusOK, gin.H{
		"schedules":   scheduleRes,
		"casts":       castRes,
		"performance": schedule_get.ParsePerformanceResponseFrom(performance),
	})
}

func Create(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	var req request
	err = c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var schedule repository.Schedule
	schedule.PerformanceID = performanceID
	schedule.Memo = req.Memo
	schedule.Open = req.Open

	if utils.CheckDateFormatInvalid(req.Date) {
		c.JSON(http.StatusBadRequest, "date format should be YYYY-MM-DD")
		return
	}

	schedule.Date = req.Date

	if req.Time != "" {
		if utils.CheckTimeFormatInvalid(req.Time) {
			c.JSON(http.StatusBadRequest, "time format should be HH:MM")
			return
		}
		schedule.Time.String = req.Time
		schedule.Time.Valid = true
	} else {
		schedule.Time.Valid = false
	}

	for _, castID := range req.Casts {
		schedule.Casts = append(schedule.Casts, repository.Cast{ID: castID})
	}

	err = repositories.schedule.Save(&schedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, schedule.UUID)
}

func Update(c *gin.Context) {
	scheduleUUID := c.Param("uuid")

	var req request
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	schedule, err := repositories.schedule.FindByUUID(scheduleUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, "schedule not found")
		return
	}

	schedule.Memo = req.Memo
	schedule.Open = req.Open

	if utils.CheckDateFormatInvalid(req.Date) {
		c.JSON(http.StatusBadRequest, "date format should be YYYY-MM-DD")
		return
	}
	schedule.Date = req.Date
	if req.Time != "" {
		if utils.CheckTimeFormatInvalid(req.Time) {
			c.JSON(http.StatusBadRequest, "time format should be HH:MM")
			return
		}
		schedule.Time.String = req.Time
		schedule.Time.Valid = true
	} else {
		schedule.Time.Valid = false
	}

	schedule.Casts = []repository.Cast{}
	for _, castID := range req.Casts {
		schedule.Casts = append(schedule.Casts, repository.Cast{ID: castID})
	}

	err = repositories.schedule.Update(schedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, schedule.UUID)
}

func Delete(c *gin.Context) {
	scheduleUUID := c.Param("uuid")

	err := repositories.schedule.Delete(scheduleUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
