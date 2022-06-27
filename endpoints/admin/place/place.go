package place

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/database"
	"github.com/kwanok/podonine/models"
	"github.com/kwanok/podonine/repository"
	"net/http"
	"strconv"
)

type Place struct {
	Title string `json:"title"`
}

var repositories Repository

type Repository struct {
	place models.PlaceRepository
}

func init() {
	repositories = Repository{
		place: &repository.PlaceRepository{Db: database.Gorm},
	}
}

func Get(c *gin.Context) {
	places := repositories.place.Get()

	if places == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, places)
}

func Find(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	place := repositories.place.Find(uint(intId))

	if place == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, place)
}

func Create(c *gin.Context) {
	var json Place
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	place := repository.Place{
		Title: json.Title,
	}

	result := repositories.place.Save(&place)

	c.JSON(http.StatusOK, result)
}
