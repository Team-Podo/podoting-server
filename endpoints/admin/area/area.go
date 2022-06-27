package area

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/database"
	"github.com/kwanok/podonine/models"
	"github.com/kwanok/podonine/repository"
	"net/http"
	"strconv"
)

type Area struct {
	Title string `json:"title"`
}

var repositories Repository

type Repository struct {
	area models.AreaRepository
}

func init() {
	repositories = Repository{
		area: &repository.AreaRepository{Db: database.Gorm},
	}
}

func Get(c *gin.Context) {
	areas := repositories.area.Get()

	if areas == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, areas)
}

func Find(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	area := repositories.area.Find(uint(intId))

	if area == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, area)
}

func Create(c *gin.Context) {
	var json Area
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	area := repository.Area{
		Title: json.Title,
	}

	result := repositories.area.Save(&area)

	c.JSON(http.StatusOK, result)
}
