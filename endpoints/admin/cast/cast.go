package cast

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/admin/cast_find"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var repositories Repository

type Repository struct {
	cast models.CastRepository
	file models.FileRepository
}

type Cast struct {
	PersonID    uint `json:"personId"`
	CharacterID uint `json:"characterId"`
}

func init() {
	repositories = Repository{
		cast: &repository.CastRepository{DB: database.Gorm},
	}
}

func Find(c *gin.Context) {
	id, err := utils.ParseUint(c.Param("cast_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(cast) id should be Integer")
		return
	}

	cast, err := repositories.cast.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "cast not found")
		return
	}

	c.JSON(http.StatusOK, cast_find.ParseResponseForm(cast))
}

func Delete(c *gin.Context) {

}
