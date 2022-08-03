package performance

import (
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
	ProductID uint   `json:"productId"`
	Title     string `json:"title"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
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

	queryMap := extractQueriesToMap(c)

	// ------ 쿼리스트링 검증 End ------

	// ------ 퍼포먼스 가져오기 Start ------

	performances := repositories.performance.GetWithQueryMap(queryMap)

	if performances == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	// ------ 퍼포먼스 가져오기 End ------

	// ------ 응답 폼 만들기 Start ------

	var performanceResponses []utils.MapSlice

	for _, performance := range performances {
		performanceResponses = append(performanceResponses, utils.BuildMapSliceByMap(map[string]any{
			"id":        performance.ID,
			"title":     performance.Title,
			"startDate": performance.StartDate,
			"endDate":   performance.EndDate,
			"product":   performance.Product,
			"schedules": performance.Schedules,
			"createdAt": performance.CreatedAt,
			"updatedAt": performance.UpdatedAt,
		}))
	}

	// ------ 응답 폼 만들기 End ------

	c.JSON(http.StatusOK, gin.H{
		"performances": performanceResponses,
		"total":        repositories.performance.GetTotalWithQueryMap(queryMap),
	})
}

func Find(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	performance := repositories.performance.FindByID(uint(intId))

	if performance == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	var performanceProduct utils.MapSlice

	if performance.Product.IsNotNil() {
		performanceProduct = utils.BuildMapSliceByMap(map[string]any{
			"id":    performance.Product.ID,
			"title": performance.Product.Title,
		})
	}

	var performanceSchedules []utils.MapSlice

	schedules := performance.Schedules
	for _, schedule := range schedules {
		performanceSchedules = append(performanceSchedules, utils.BuildMapSliceByMap(map[string]any{
			"uuid": schedule.UUID,
			"memo": schedule.Memo,
			"date": schedule.Date,
			"time": schedule.Time,
		}))
	}

	result := utils.BuildMapSliceByMap(map[string]any{
		"id":        performance.ID,
		"title":     performance.Title,
		"startDate": performance.StartDate,
		"endDate":   performance.EndDate,
		"product":   performanceProduct,
		"schedules": performanceSchedules,
		"createdAt": performance.CreatedAt,
		"updatedAt": performance.UpdatedAt,
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

	performance := repository.Performance{
		Product:   &repository.Product{ID: json.ProductID},
		Title:     json.Title,
		StartDate: json.StartDate,
		EndDate:   json.EndDate,
	}

	err := repositories.performance.Save(&performance)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, performance.ID)
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

	performance := repository.Performance{
		ID:        uint(intId),
		Product:   &repository.Product{ID: json.ProductID},
		Title:     json.Title,
		StartDate: json.StartDate,
		EndDate:   json.EndDate,
	}

	err = repositories.performance.Update(&performance)

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, performance.ID)
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	err = repositories.performance.Delete(uint(intId))
	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, nil)
}

func extractQueriesToMap(c *gin.Context) map[string]any {
	query := map[string]any{
		"limit":    getLimitQuery(c),
		"offset":   getOffsetQuery(c),
		"reversed": getReverseQuery(c),
	}

	return query
}

func getLimitQuery(c *gin.Context) int {
	limitQuery := c.Query("limit")
	if limitQuery == "" {
		return 10
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		return 10
	}

	return limit
}

func getOffsetQuery(c *gin.Context) int {
	offsetQuery := c.Query("offset")
	if offsetQuery == "" {
		return 0
	}

	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		return 0
	}

	return offset
}

func getReverseQuery(c *gin.Context) bool {
	reversedQuery := c.Query("reversed")
	if reversedQuery == "" {
		return false
	}

	reversed, err := strconv.ParseBool(reversedQuery)
	if err != nil {
		return false
	}

	return reversed
}
