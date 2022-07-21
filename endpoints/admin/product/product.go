package product

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/database"
	"github.com/kwanok/podonine/models"
	"github.com/kwanok/podonine/repository"
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

	products := repositories.product.Get(map[string]any{
		"limit":    limit,
		"offset":   offset,
		"reversed": reversed,
	})

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, products)
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
