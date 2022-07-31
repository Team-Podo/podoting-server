package endpoints

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/gin-gonic/gin"
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
		product: &repository.ProductRepository{Db: database.Gorm},
	}
}

func GetProduct(c *gin.Context) {
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

	result := repositories.product.Save(&product)

	c.JSON(http.StatusOK, result)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	repositories.product.Delete(uint(intId))
}
