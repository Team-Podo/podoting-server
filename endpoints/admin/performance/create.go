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
		MainAreaID:  &json.MainAreaID,
		Title:       json.Title,
		RunningTime: json.RunningTime,
		StartDate:   json.StartDate,
		EndDate:     json.EndDate,
		Rating:      json.Rating,
	}

	ab, err := repositories.areaBoilerplate.GetAreaBoilerplateByAreaID(*performance.MainAreaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "mainAreaID is not valid")
		return
	}

	err = repositories.performance.Save(&performance)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "database_error: performance save failed")
		return
	}

	err = repositories.areaBoilerplate.SaveSeats(ab, performance.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database_error: seats save failed")
		return
	}

	c.JSON(http.StatusCreated, performance.ID)
}

func UploadThumbnailImage(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	performance := repositories.performance.FindByID(performanceID)

	if performance == nil {
		c.JSON(http.StatusNotFound, "performance not found please check id")
		return
	}

	thumbnailImage, fileHeader, err := c.Request.FormFile("thumbnailImage")
	if err != nil {
		c.JSON(http.StatusBadRequest, "thumbnailImage is required")
		return
	}

	if !utils.CheckFileExtension(fileHeader) {
		c.JSON(http.StatusBadRequest, "file extension is not allowed")
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)
	filePath := fmt.Sprintf("performance/%d/thumbnail", performanceID)
	file, err := aws.S3.UploadFile(thumbnailImage, filePath, fileExtension)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "file upload failed")
		return
	}

	performance.Thumbnail = &repository.File{
		Path: *file.Key,
		Size: fileHeader.Size,
	}

	err = repositories.file.Save(performance.Thumbnail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database_error: file db save failed")
		return
	}

	performance.ThumbnailID = &performance.Thumbnail.ID

	err = repositories.performance.Update(performance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database_error: performance update failed")
		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"thumbnail": performance.Thumbnail.FullPath(),
	})
}
