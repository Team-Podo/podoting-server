package place

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
	file  models.FileRepository
	place models.PlaceRepository
}

func init() {
	repositories = Repository{
		file:  &repository.FileRepository{DB: database.Gorm},
		place: &repository.PlaceRepository{DB: database.Gorm},
	}
}

func UploadPlaceImage(c *gin.Context) {
	performanceID, err := getIntParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	placeID, err := getIntParam(c, "place_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, "cast_id should be Integer")
		return
	}

	place, err := repositories.place.FindByID(uint(placeID))

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	mainImage, fileHeader, err := c.Request.FormFile("placeImage")

	if err != nil {
		c.JSON(http.StatusBadRequest, "profileImage is required")
		return
	}

	if !utils.CheckFileExtension(fileHeader) {
		c.JSON(http.StatusBadRequest, "File extension is not allowed")
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)

	filePath := fmt.Sprintf("/performances/%d/places/%d/place-images", performanceID, placeID)

	file, err := aws.S3.UploadFile(mainImage, filePath, fileExtension)

	if err != nil {
		c.JSON(http.StatusBadRequest, "mainImage should be File")
		return
	}

	place.PlaceImage = &repository.File{
		Path: *file.Key,
		Size: fileHeader.Size,
	}

	err = repositories.file.Save(place.PlaceImage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	err = repositories.place.Update(place)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"profileImage": fmt.Sprintf("%s/%s", os.Getenv("CDN_URL"), place.PlaceImage.Path),
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
