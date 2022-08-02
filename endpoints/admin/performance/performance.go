package performance

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/endpoints/admin/product"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var repositories Repository

type request struct {
	ProductID uint   `json:"productId"`
	Title     string `json:"title"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type Performance struct {
	Id        uint
	Product   models.Product
	Title     string
	StartDate string
	EndDate   string
}

func (p *Performance) GetCreatedAt() string {
	return ""
}

func (p *Performance) GetUpdatedAt() string {
	return ""
}

func (p *Performance) GetId() uint {
	return p.Id
}

func (p *Performance) GetProduct() models.Product {
	return p.Product
}

func (p *Performance) GetSchedules() []models.Schedule {
	return nil
}

func (p *Performance) GetTitle() string {
	return p.Title
}

func (p *Performance) GetStartDate() string {
	return p.StartDate
}

func (p *Performance) GetEndDate() string {
	return p.EndDate
}

type Repository struct {
	performance models.PerformanceRepository
}

func init() {
	repositories = Repository{
		performance: &repository.PerformanceRepository{Db: database.Gorm},
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

	performances := repositories.performance.Get(query)

	if performances == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	// ------ 퍼포먼스 가져오기 End ------

	// ------ 응답 폼 만들기 Start ------

	var performanceResponses []utils.MapSlice

	for _, performance := range performances {
		performanceResponses = append(performanceResponses, utils.MapSlice{
			utils.MapItem{Key: "id", Value: performance.GetId()},
			utils.MapItem{Key: "title", Value: performance.GetTitle()},
			utils.MapItem{Key: "startDate", Value: performance.GetStartDate()},
			utils.MapItem{Key: "endDate", Value: performance.GetEndDate()},
			utils.MapItem{Key: "createdAt", Value: performance.GetCreatedAt()},
			utils.MapItem{Key: "updatedAt", Value: performance.GetUpdatedAt()},
		})
	}

	// ------ 응답 폼 만들기 End ------

	c.JSON(http.StatusOK, gin.H{
		"performances": performanceResponses,
		"total":        repositories.performance.GetTotal(query),
	})
}

func Find(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	performance := repositories.performance.Find(uint(intId))

	if performance == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	var performanceProduct *utils.MapSlice

	if performance.GetProduct().IsNotNil() {
		performanceProduct = &utils.MapSlice{
			utils.MapItem{Key: "id", Value: performance.GetProduct().GetId()},
			utils.MapItem{Key: "title", Value: performance.GetProduct().GetTitle()},
		}
	}

	var performanceSchedules []*utils.MapSlice

	if performance.GetProduct().IsNotNil() {
		schedules := performance.GetSchedules()
		for _, schedule := range schedules {
			performanceSchedules = append(performanceSchedules, &utils.MapSlice{
				utils.MapItem{Key: "uuid", Value: schedule.GetUUID()},
				utils.MapItem{Key: "memo", Value: schedule.GetMemo()},
				utils.MapItem{Key: "date", Value: schedule.GetDate()},
				utils.MapItem{Key: "time", Value: schedule.GetTime()},
			})
		}
	}

	result := utils.BuildMapSliceByMap(map[string]any{
		"id":        performance.GetId(),
		"title":     performance.GetTitle(),
		"startDate": performance.GetStartDate(),
		"endDate":   performance.GetEndDate(),
		"product":   performanceProduct,
		"schedules": performanceSchedules,
		"createdAt": performance.GetCreatedAt(),
		"updatedAt": performance.GetUpdatedAt(),
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

	performance := repositories.performance.Save(&Performance{
		Product:   &product.Product{ID: json.ProductID},
		Title:     json.Title,
		StartDate: json.StartDate,
		EndDate:   json.EndDate,
	})

	if performance == nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, performance.GetId())
}

func Update(c *gin.Context) {
	var json request

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	performance := repositories.performance.Update(&Performance{
		Id:        uint(intId),
		Product:   &product.Product{ID: json.ProductID},
		Title:     json.Title,
		StartDate: json.StartDate,
		EndDate:   json.EndDate,
	})

	if performance == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, performance.GetId())
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	repositories.performance.Delete(uint(intId))

	c.JSON(http.StatusOK, nil)
}
