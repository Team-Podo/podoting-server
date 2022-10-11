package content

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Repository struct {
	content models.PerformanceContentRepository
}

type Request struct {
	Content       string `json:"content" binding:"required"`
	ManagingTitle string `json:"managingTitle" binding:"required"`
	Visible       bool   `json:"visible"`
}

var repositories Repository

func init() {
	repositories = Repository{
		content: &repository.PerformanceContentRepository{DB: database.Gorm},
	}
}

func Create(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	fmt.Println(performanceID)

	var request Request
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	content := repository.PerformanceContent{
		PerformanceID: performanceID,
		Content:       request.Content,
		ManagingTitle: request.ManagingTitle,
	}

	err = repositories.content.Save(&content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func Update(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	contentID, err := utils.ParseUint(c.Param("content_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(content) id should be Integer")
		return
	}

	var request Request
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	content := repositories.content.FindOneByID(contentID)
	if content == nil {
		c.JSON(http.StatusNotFound, "content not found")
		return
	}

	if content.PerformanceID != performanceID {
		c.JSON(http.StatusForbidden, "content not found")
		return
	}

	content.Content = request.Content
	content.ManagingTitle = request.ManagingTitle
	content.Visible = request.Visible

	err = repositories.content.Save(content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
