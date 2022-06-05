package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/models"
	"github.com/kwanok/podonine/repository"
	"net/http"
	"strconv"
)

type Product struct {
	Title string `json:"title"`
}

var repositories Repository

type Repository struct {
	product models.ProductRepository
}

func init() {
	repositories = Repository{
		product: &repository.ProductRepository{Db: repository.Gorm},
	}
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	product := repositories.product.GetProductById(uint(intId))

	if product == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, product)
}

func SaveProduct(c *gin.Context) {
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

	result := repositories.product.SaveProduct(&product)

	c.JSON(http.StatusOK, result)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	repositories.product.DeleteProductById(uint(intId))
}
