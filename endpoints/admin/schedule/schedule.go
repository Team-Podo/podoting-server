package schedule

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance"
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

type Schedule struct {
	Id          uint
	UUID        string
	Performance performance.Performance
	Memo        string
	Date        string
	Time        string
}

func (s *Schedule) GetUUID() string {
	return s.UUID
}

func (s *Schedule) GetPerformance() models.Performance {
	return nil
}

func (s *Schedule) GetMemo() string {
	return s.Memo
}

func (s *Schedule) GetDate() string {
	return s.Date
}

func (s *Schedule) GetTime() string {
	return s.Time
}

func (s *Schedule) GetCreatedAt() string {
	return ""
}

func (s *Schedule) GetUpdatedAt() string {
	return ""
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
		scheduleResponses = append(scheduleResponses, utils.MapSlice{
			utils.MapItem{Key: "uuid", Value: schedule.GetUUID()},
			utils.MapItem{Key: "performance", Value: schedule.GetPerformance()},
			utils.MapItem{Key: "memo", Value: schedule.GetMemo()},
			utils.MapItem{Key: "date", Value: schedule.GetDate()},
			utils.MapItem{Key: "time", Value: schedule.GetTime()},
			utils.MapItem{Key: "createdAt", Value: schedule.GetCreatedAt()},
			utils.MapItem{Key: "updatedAt", Value: schedule.GetUpdatedAt()},
		})
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

	c.JSON(http.StatusOK, utils.MapSlice{
		utils.MapItem{Key: "uuid", Value: schedule.GetUUID()},
		utils.MapItem{Key: "performance", Value: schedule.GetPerformance()},
		utils.MapItem{Key: "memo", Value: schedule.GetMemo()},
		utils.MapItem{Key: "date", Value: schedule.GetDate()},
		utils.MapItem{Key: "time", Value: schedule.GetTime()},
		utils.MapItem{Key: "createdAt", Value: schedule.GetCreatedAt()},
		utils.MapItem{Key: "updatedAt", Value: schedule.GetUpdatedAt()},
	})
}

func Create(c *gin.Context) {
	var json request
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	schedule := repositories.schedule.Save(&Schedule{})

	if schedule == nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, schedule.GetUUID())
}

func CreateMany(c *gin.Context) {
	// TODO endpoint 구현 필요
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

	schedule := repositories.schedule.Update(&Schedule{
		UUID: uuid,
		Memo: json.Memo,
		Date: json.Date,
		Time: json.Time,
	})

	if schedule == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, schedule.GetUUID())
}

func Delete(c *gin.Context) {
	uuid := c.Param("uuid")

	repositories.schedule.Delete(uuid)

	c.JSON(http.StatusOK, nil)
}
