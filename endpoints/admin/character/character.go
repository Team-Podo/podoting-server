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

type request struct {
	Name string `json:"name" binding:"required"`
}

var repositories Repository

func init() {
	repositories = Repository{
		character: &repository.CharacterRepository{DB: database.Gorm},
	}
}

func Delete(c *gin.Context) {
	characterID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(character) id should be Integer")
		return
	}

	err = repositories.character.Delete(characterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func Update(c *gin.Context) {
	characterID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(character) id should be Integer")
		return
	}

	var req request
	err = c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var character repository.Character
	character.ID = characterID
	character.Name = req.Name

	err = repositories.character.Update(&character)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, character.ID)
}
