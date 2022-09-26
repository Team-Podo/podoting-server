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

func Find(c *gin.Context) {
	placeID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	place, err := repositories.place.FindByID(placeID)
	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, place)
}

func Get(c *gin.Context) {
	place, err := repositories.place.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, place)
}

func Create(c *gin.Context) {
	var place repository.Place

	err := c.BindJSON(&place)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON")
		return
	}

	err = repositories.place.Create(&place)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusCreated, place.ID)
}

func Update(c *gin.Context) {
	placeID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	var request repository.Place
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON")
		return
	}

	place, err := repositories.place.FindByID(placeID)
	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	place.Name = request.Name

	err = repositories.place.Update(place)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, place.ID)
}

func Delete(c *gin.Context) {
	placeID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	place, err := repositories.place.FindByID(placeID)
	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	err = repositories.place.Delete(place.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, place.ID)
}

func UploadPlaceImage(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	placeID, err := utils.ParseUint(c.Param("place_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "place_id should be Integer")
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
