package performance

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/database"
	"github.com/kwanok/podonine/endpoints/admin/product"
	"github.com/kwanok/podonine/models"
	"github.com/kwanok/podonine/repository"
	"net/http"
	"strconv"
)

var repositories Repository

type Request struct {
	ProductID uint   `json:"ProductID"`
	Title     string `json:"Title"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
}

type Performance struct {
	Id        uint
	Product   models.Product
	Title     string
	StartDate string
	EndDate   string
}

func (p *Performance) GetId() uint {
	return p.Id
}

func (p *Performance) GetProduct() models.Product {
	return p.Product
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

	// ------ 쿼리스트링 검증 End ------

	// ------ 퍼포먼스 가져오기 Start ------

	performances := repositories.performance.Get(map[string]any{
		"limit":    limit,
		"offset":   offset,
		"reversed": reversed,
	})

	if performances == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	// ------ 퍼포먼스 가져오기 End ------

	c.JSON(http.StatusOK, performances)
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

	c.JSON(http.StatusOK, performance)
}

func Create(c *gin.Context) {
	var json Request
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
	var json Request

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
