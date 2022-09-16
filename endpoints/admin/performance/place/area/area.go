package area

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
	file models.FileRepository
	area models.AreaRepository
}

func init() {
	repositories = Repository{
		file: &repository.FileRepository{DB: database.Gorm},
		area: &repository.AreaRepository{DB: database.Gorm},
	}
}

func UploadAreaImage(c *gin.Context) {
	performanceID, err := getIntParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	placeID, err := getIntParam(c, "place_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, "place_id should be Integer")
		return
	}

	areaID, err := getIntParam(c, "area_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, "area_id should be Integer")
		return
	}

	area := repositories.area.FindOne(uint(areaID))

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	mainImage, fileHeader, err := c.Request.FormFile("backgroundImage")

	if err != nil {
		c.JSON(http.StatusBadRequest, "backgroundImage is required")
		return
	}

	if !utils.CheckFileExtension(fileHeader) {
		c.JSON(http.StatusBadRequest, "File extension is not allowed")
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)

	filePath := fmt.Sprintf("/performances/%d/places/%d/areas/%d/background-images", performanceID, placeID, areaID)

	file, err := aws.S3.UploadFile(mainImage, filePath, fileExtension)

	if err != nil {
		c.JSON(http.StatusBadRequest, "mainImage should be File")
		return
	}

	area.BackgroundImage = &repository.File{
		Path: *file.Key,
		Size: fileHeader.Size,
	}

	err = repositories.file.Save(area.BackgroundImage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	err = repositories.area.Update(area)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"profileImage": fmt.Sprintf("%s/%s", os.Getenv("CDN_URL"), area.BackgroundImage.Path),
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
