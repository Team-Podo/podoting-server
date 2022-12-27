package cast

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
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
		if cast.ID != 0 {
			nc, castErr := repositories.cast.FindOneByID(cast.ID)
			if castErr != nil {
				c.JSON(http.StatusNotFound, "cast not found")
				return
			}

			nc.PerformanceID = performanceID
			nc.PersonID = cast.PersonID
			nc.CharacterID = cast.CharacterID

			newCasts = append(newCasts, *nc)
		} else {
			var newCast *repository.Cast

			newCast.PerformanceID = performanceID
			newCast.PersonID = cast.PersonID
			newCast.CharacterID = cast.CharacterID

			newCasts = append(newCasts, *newCast)
		}

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
