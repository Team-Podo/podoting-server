package cast

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/admin/cast_get"
	"github.com/Team-Podo/podoting-server/response/admin/cast_save"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var repositories Repository

type Repository struct {
	cast models.CastRepository
	file models.FileRepository
}

type Cast struct {
	ID          uint `json:"id"`
	PersonID    uint `json:"personID" binding:"required"`
	CharacterID uint `json:"characterID" binding:"required"`
}

func init() {
	repositories = Repository{
		cast: &repository.CastRepository{DB: database.Gorm},
		file: &repository.FileRepository{DB: database.Gorm},
	}
}

func Get(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	casts, err := repositories.cast.FindByPerformanceID(performanceID)
	if err != nil {
		c.JSON(http.StatusNotFound, "casts not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"casts": cast_get.ParseResponseForm(casts),
	})
}

func CreateMany(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	var casts []Cast
	err = c.BindJSON(&casts)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON")
		return
	}

	var newCasts []repository.Cast

	for _, cast := range casts {
		newCast := repository.Cast{
			ID:            cast.ID,
			PerformanceID: performanceID,
			PersonID:      cast.PersonID,
			CharacterID:   cast.CharacterID,
		}
		newCasts = append(newCasts, newCast)
	}

	err = repositories.cast.CreateMany(newCasts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "database error: cast create failed")
		log.Fatal(err)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, "database error: performance cast create failed")
		log.Fatal(err)
		return
	}

	response := cast_save.ParseResponseForm(newCasts)

	c.JSON(http.StatusOK, gin.H{
		"casts": response,
	})
}
