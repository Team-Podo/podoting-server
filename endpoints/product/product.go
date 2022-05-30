package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/repository"
	"net/http"
	"strconv"
)

type Product struct {
	Title string `json:"title"`
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		c.Abort()
		return
	}

	productRepository := repository.ProductRepository{Db: repository.Gorm}
	product := productRepository.GetProductById(uint(intId))

	if product == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		c.Abort()
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

	product := repository.Product{Title: json.Title}

	productRepository := repository.ProductRepository{Db: repository.Gorm}
	productRepository.SaveProduct(&product)

	c.JSON(http.StatusOK, "OK")
}
