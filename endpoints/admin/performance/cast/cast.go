package cast

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/Team-Podo/podoting-server/utils/aws"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var repositories Repository

type Repository struct {
	cast models.CastRepository
	file models.FileRepository
}

func init() {
	repositories = Repository{
		cast: &repository.CastRepository{DB: database.Gorm},
		file: &repository.FileRepository{DB: database.Gorm},
	}
}

func UploadProfileImage(c *gin.Context) {
	performanceID, err := getIntParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	castID, err := getIntParam(c, "cast_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, "cast_id should be Integer")
		return
	}

	cast, err := repositories.cast.FindByID(uint(castID))

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	mainImage, fileHeader, err := c.Request.FormFile("profileImage")

	if err != nil {
		c.JSON(http.StatusBadRequest, "profileImage is required")
		return
	}

	if !utils.CheckFileExtension(fileHeader) {
		c.JSON(http.StatusBadRequest, "File extension is not allowed")
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)

	filePath := fmt.Sprintf("/performances/%d/casts/%d/profile-images", performanceID, castID)

	file, err := aws.S3.UploadFile(mainImage, filePath, fileExtension)

	if err != nil {
		c.JSON(http.StatusBadRequest, "mainImage should be File")
		return
	}

	cast.ProfileImage = &repository.File{
		Path: *file.Key,
		Size: fileHeader.Size,
	}

	err = repositories.file.Save(cast.ProfileImage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	err = repositories.cast.Update(cast)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"profileImage": fmt.Sprintf("%s/%s", os.Getenv("CDN_URL"), cast.ProfileImage.Path),
	})
}

func getIntParam(c *gin.Context, param string) (int, error) {
	id := c.Param(param)

	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return intId, nil
}
