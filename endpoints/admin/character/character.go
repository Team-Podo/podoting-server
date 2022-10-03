package character

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Repository struct {
	character models.CharacterRepository
}

var repositories Repository

func init() {
	repositories = Repository{
		character: &repository.CharacterRepository{DB: database.Gorm},
	}
}

func Delete(c *gin.Context) {
	characterID, err := utils.ParseUint(c.Param("character_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(character) id should be Integer")
		return
	}

	err = repositories.character.Delete(characterID)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
