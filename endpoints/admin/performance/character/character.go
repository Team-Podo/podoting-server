package character

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
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

func Create(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("performance_id"))
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

	character := repositories.character.Create(&repository.Character{
		PerformanceID: performanceID,
		Name:          req.Name,
	})

	c.JSON(http.StatusOK, character)
}

func Update(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("performance_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	characterID, err := utils.ParseUint(c.Param("character_id"))
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

	character := repositories.character.Update(&repository.Character{
		PerformanceID: performanceID,
		ID:            characterID,
		Name:          req.Name,
	})

	c.JSON(http.StatusOK, character)
}
