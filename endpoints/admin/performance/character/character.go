package character

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/admin/character_get"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type request struct {
	Name string `json:"name" binding:"required"`
}

type Repository struct {
	character models.CharacterRepository
}

var repositories Repository

func init() {
	repositories = Repository{
		character: &repository.CharacterRepository{DB: database.Gorm},
	}
}

func Get(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	characters, err := repositories.character.FindByPerformanceID(performanceID)
	if err != nil {
		c.JSON(http.StatusNotFound, "character not found")
		return
	}

	c.JSON(http.StatusOK, character_get.ParseResponseForm(characters))
}

func Create(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	var req request
	err = c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var character repository.Character
	character.Name = req.Name
	character.PerformanceID = performanceID

	err = repositories.character.Create(&character)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, character.ID)
}
