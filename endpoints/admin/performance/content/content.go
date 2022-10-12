package content

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/admin/performance_content_find"
	"github.com/Team-Podo/podoting-server/response/admin/performance_content_get"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/Team-Podo/podoting-server/utils/aws"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type Repository struct {
	content models.PerformanceContentRepository
	file    models.FileRepository
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
		file:    &repository.FileRepository{DB: database.Gorm},
	}
}

func Find(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	contents := repositories.content.FindByPerformanceID(performanceID)

	c.JSON(http.StatusOK, gin.H{
		"contents": performance_content_get.ParseResponse(contents),
	})
}

func FindOne(c *gin.Context) {
	contentID, err := utils.ParseUint(c.Param("content_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(content) id should be Integer")
		return
	}

	content := repositories.content.FindOneByID(contentID)
	if content == nil {
		c.JSON(http.StatusNotFound, "content not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": performance_content_find.ParseResponse(content),
	})
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

func UploadContentImage(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	contentImage, fileHeader, err := c.Request.FormFile("contentImage")
	if err != nil {
		c.JSON(http.StatusBadRequest, "contentImage is required")
		return
	}

	if !utils.CheckFileExtension(fileHeader) {
		c.JSON(http.StatusBadRequest, "file extension is not allowed")
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)
	filePath := fmt.Sprintf("performance/%d/contents", performanceID)
	file, err := aws.S3.UploadFile(contentImage, filePath, fileExtension)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "file upload failed")
		return
	}

	newFile := &repository.File{
		Path: *file.Key,
		Size: fileHeader.Size,
	}

	err = repositories.file.Save(newFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database_error: file db save failed")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": newFile.FullPath(),
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

func Delete(c *gin.Context) {
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

	content := repositories.content.FindOneByID(contentID)
	if content == nil {
		c.JSON(http.StatusNotFound, "content not found")
		return
	}

	if content.PerformanceID != performanceID {
		c.JSON(http.StatusForbidden, "content not found")
		return
	}

	err = repositories.content.Delete(content.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
