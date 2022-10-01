package area

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	FindResponse "github.com/Team-Podo/podoting-server/response/admin/area_find"
	GetResponse "github.com/Team-Podo/podoting-server/response/admin/area_get"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/Team-Podo/podoting-server/utils/aws"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

var repositories Repository

type Repository struct {
	file models.FileRepository
	area models.AreaRepository
}

type request struct {
	Name string `json:"name" binding:"required"`
}

func init() {
	repositories = Repository{
		file: &repository.FileRepository{DB: database.Gorm},
		area: &repository.AreaRepository{DB: database.Gorm},
	}
}

func Get(c *gin.Context) {
	placeID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(place) id should be Integer")
		return
	}

	areas := repositories.area.GetByPlaceID(placeID)
	if len(areas) == 0 {
		c.JSON(http.StatusNotFound, "areas are not found")
		return
	}

	c.JSON(http.StatusOK, GetResponse.ParseResponseForm(areas))
}

func Find(c *gin.Context) {
	placeID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(place) id should be Integer")
		return
	}

	areaID, err := utils.ParseUint(c.Param("area_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(area) id should be Integer")
		return
	}

	area := repositories.area.FindOne(placeID, areaID)

	if err != nil {
		c.JSON(http.StatusNotFound, "area is not found")
		return
	}

	c.JSON(http.StatusOK, FindResponse.ParseResponseForm(area))
}

func Create(c *gin.Context) {
	placeID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(place) id should be Integer")
		return
	}

	var req request
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "name is required")
		return
	}

	area := &repository.Area{
		PlaceID: uint(placeID),
		Name:    req.Name,
	}

	err = repositories.area.Create(area)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database error: area create failed")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, area.ID)
}

func UploadAreaImage(c *gin.Context) {
	placeID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(place) id should be Integer")
		return
	}

	areaID, err := utils.ParseUint(c.Param("area_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(area) id should be Integer")
		return
	}

	area := repositories.area.FindOne(placeID, areaID)

	if err != nil {
		c.JSON(http.StatusNotFound, "area is not found")
		return
	}

	mainImage, fileHeader, err := c.Request.FormFile("backgroundImage")

	if err != nil {
		c.JSON(http.StatusBadRequest, "backgroundImage is required")
		return
	}

	if !utils.CheckFileExtension(fileHeader) {
		c.JSON(http.StatusBadRequest, "file extension is not allowed")
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)

	filePath := fmt.Sprintf("/places/%d/areas/%d/background-images", placeID, areaID)

	file, err := aws.S3.UploadFile(mainImage, filePath, fileExtension)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "file upload failed")
		return
	}

	area.BackgroundImage = &repository.File{
		Path: *file.Key,
		Size: fileHeader.Size,
	}

	err = repositories.file.Save(area.BackgroundImage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "database_error: file save failed")
		log.Fatal(err)
		return
	}

	err = repositories.area.Update(area)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "database_error: area update failed")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"backgroundImage": area.BackgroundImage.FullPath(),
	})
}

func Update(c *gin.Context) {
	placeID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(place) id should be Integer")
		return
	}

	areaID, err := utils.ParseUint(c.Param("area_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(area) id should be Integer")
		return
	}

	var req request
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "name is required")
		return
	}

	area := repositories.area.FindOne(placeID, areaID)
	if area == nil {
		c.JSON(http.StatusNotFound, "area is not found please check id")
		return
	}

	area.Name = req.Name

	err = repositories.area.Update(area)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database error: area update failed")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, area.ID)
}

func Delete(c *gin.Context) {
	placeID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(place) id should be Integer")
		return
	}

	areaID, err := utils.ParseUint(c.Param("area_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(area) id should be Integer")
		return
	}

	area := repositories.area.FindOne(placeID, areaID)
	if area == nil {
		c.JSON(http.StatusNotFound, "area is not found please check id")
		return
	}

	err = repositories.area.Delete(area)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database error: area delete failed")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
