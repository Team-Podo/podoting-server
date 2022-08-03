package schedule

import (
	"database/sql"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var repositories Repository

type request struct {
	PerformanceId string `json:"performanceId"`
	Memo          string `json:"memo"`
	Date          string `json:"date"`
	Time          string `json:"time"`
}

type Repository struct {
	schedule models.ScheduleRepository
}

func init() {
	repositories = Repository{
		schedule: &repository.ScheduleRepository{Db: database.Gorm},
	}
}

func Get(c *gin.Context) {
	// ------ 쿼리스트링 검증 Start ------

	limitQuery := c.Query("limit")
	offsetQuery := c.Query("offset")
	reversedQuery := c.Query("reversed")

	var limit int
	var offset int
	var reversed = false
	var err error

	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, "limit should be Integer")
			return
		}
	}

	if offsetQuery != "" {
		offset, err = strconv.Atoi(offsetQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, "offset should be Integer")
			return
		}
	}

	if reversedQuery != "" {
		reversed = true
	}

	query := map[string]any{
		"limit":    limit,
		"offset":   offset,
		"reversed": reversed,
	}

	// ------ 쿼리스트링 검증 End ------

	// ------ 퍼포먼스 가져오기 Start ------

	schedules := repositories.schedule.Get(query)

	if schedules == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	// ------ 스케줄 가져오기 End ------

	// ------ 응답 폼 만들기 Start ------

	var scheduleResponses []utils.MapSlice

	for _, schedule := range schedules {
		scheduleResponses = append(scheduleResponses, utils.BuildMapSliceByMap(map[string]any{
			"uuid":      schedule.UUID,
			"memo":      schedule.Memo,
			"date":      schedule.Date,
			"time":      schedule.Time,
			"createdAt": schedule.CreatedAt,
			"updatedAt": schedule.UpdatedAt,
		}))
	}

	// ------ 응답 폼 만들기 End ------

	c.JSON(http.StatusOK, gin.H{
		"schedules": scheduleResponses,
		"total":     repositories.schedule.GetTotal(query),
	})
}

func Find(c *gin.Context) {
	uuid := c.Param("uuid")

	schedule := repositories.schedule.Find(uuid)

	if schedule == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	result := utils.BuildMapSliceByMap(map[string]any{
		"uuid":        schedule.UUID,
		"performance": schedule.Performance,
		"memo":        schedule.Memo,
		"date":        schedule.Date,
		"time":        schedule.Time,
		"createdAt":   schedule.CreatedAt,
		"updatedAt":   schedule.UpdatedAt,
	})

	c.JSON(http.StatusOK, result)
}

func Create(c *gin.Context) {
	var json request
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var schedule repository.Schedule

	err := repositories.schedule.Save(&schedule)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, schedule)
}

func Update(c *gin.Context) {
	var json request

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	uuid := c.Param("uuid")

	schedule := repository.Schedule{
		UUID: uuid,
		Memo: json.Memo,
		Date: json.Date,
		Time: sql.NullString{String: json.Time},
	}

	err := repositories.schedule.Update(&schedule)

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, schedule)
}

func Delete(c *gin.Context) {
	uuid := c.Param("uuid")

	err := repositories.schedule.Delete(uuid)

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
	}

	c.JSON(http.StatusOK, nil)
}
