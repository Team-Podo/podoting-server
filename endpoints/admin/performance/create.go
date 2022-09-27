package performance

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/Team-Podo/podoting-server/utils/aws"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func Create(c *gin.Context) {
	var json request
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	performance := repository.Performance{
		PlaceID:     &json.PlaceID,
		Title:       json.Title,
		RunningTime: json.RunningTime,
		StartDate:   json.StartDate,
		EndDate:     json.EndDate,
		Rating:      json.Rating,
	}

	if json.ProductID != 0 {
		performance.ProductID = &json.ProductID
	}

	err := repositories.performance.Save(&performance)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, performance.ID)
}

func UploadThumbnailImage(c *gin.Context) {
	performanceID, err := parseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	performance := repositories.performance.FindByID(performanceID)

	if performance == nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	thumbnailImage, fileHeader, err := c.Request.FormFile("thumbnailImage")
	if err != nil {
		c.JSON(http.StatusBadRequest, "mainImage is required")
		return
	}

	if !utils.CheckFileExtension(fileHeader) {
		c.JSON(http.StatusBadRequest, "File extension is not allowed")
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)
	filePath := fmt.Sprintf("performance/%d/thumbnail", performanceID)
	file, err := aws.S3.UploadFile(thumbnailImage, filePath, fileExtension)

	if err != nil {
		c.JSON(http.StatusBadRequest, "mainImage should be File")
		return
	}

	performance.Thumbnail = &repository.File{
		Path: *file.Key,
		Size: fileHeader.Size,
	}

	err = repositories.file.Save(performance.Thumbnail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "File Save Error")
		return
	}

	performance.ThumbnailID = &performance.Thumbnail.ID

	err = repositories.performance.Update(performance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Performance Update Error")
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"thumbnail": performance.Thumbnail.FullPath(),
	})
}
