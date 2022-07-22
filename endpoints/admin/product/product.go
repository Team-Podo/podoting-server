package product

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/database"
	"github.com/kwanok/podonine/models"
	"github.com/kwanok/podonine/repository"
	"github.com/kwanok/podonine/utils"
	"net/http"
	"strconv"
)

type Product struct {
	ID    uint
	Title string `json:"title"`
}

func (p *Product) GetId() uint {
	return p.ID
}

func (p *Product) GetTitle() string {
	return p.Title
}

func (p *Product) GetPlace() models.Place {
	return nil
}

func (p *Product) GetCreatedAt() string {
	return ""
}

func (p *Product) GetUpdatedAt() string {
	return ""
}

func (p *Product) IsNil() bool {
	if p == nil {
		return true
	}

	return false
}

func (p *Product) IsNotNil() bool {
	if p == nil {
		return false
	}

	return true
}

var repositories Repository

type Repository struct {
	product models.ProductRepository
}

func init() {
	repositories = Repository{
		product: &repository.ProductRepository{Db: database.Gorm},
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

	// ------ 상품 가져오기 Start ------

	products := repositories.product.Get(query)

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	// ------ 상품 가져오기 End ------

	// ------ 응답 폼 만들기 Start ------

	var productResponses []utils.MapSlice

	for _, product := range products {
		productResponses = append(productResponses, utils.MapSlice{
			utils.MapItem{Key: "id", Value: product.GetId()},
			utils.MapItem{Key: "title", Value: product.GetTitle()},
			utils.MapItem{Key: "createdAt", Value: product.GetCreatedAt()},
			utils.MapItem{Key: "updatedAt", Value: product.GetUpdatedAt()},
		})
	}

	// ------ 응답 폼 만들기 End ------

	c.JSON(http.StatusOK, map[string]any{
		"products": products,
	})
}

func Find(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	product := repositories.product.Find(uint(intId))

	if product == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, product)
}

func Create(c *gin.Context) {
	var json Product
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	product := repository.Product{
		Title: json.Title,
	}

	result := repositories.product.Save(&product)

	c.JSON(http.StatusOK, result)
}
