package cast

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/admin/cast_get"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/Team-Podo/podoting-server/utils/aws"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

var repositories Repository

type Repository struct {
	cast models.CastRepository
	file models.FileRepository
}

type Cast struct {
	PersonID    uint `json:"personId"`
	CharacterID uint `json:"characterId"`
}

func init() {
	repositories = Repository{
		cast: &repository.CastRepository{DB: database.Gorm},
		file: &repository.FileRepository{DB: database.Gorm},
	}
}

func Get(c *gin.Context) {
	performanceID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(performance) id should be Integer")
		return
	}

	casts, err := repositories.cast.GetByPerformanceID(performanceID)
	if err != nil {
		c.JSON(http.StatusNotFound, "casts not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"casts": cast_get.ParseResponseForm(casts),
	})
}

func Create(c *gin.Context) {
	var cast Cast
	err := c.BindJSON(&cast)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON")
		return
	}

	newCast := repository.Cast{
		PersonID:    cast.PersonID,
		CharacterID: cast.CharacterID,
	}

	err = repositories.cast.Create(&newCast)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "database error: cast create failed")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cast": newCast,
	})
}

func UploadProfileImage(c *gin.Context) {
	castID, err := utils.ParseUint(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "(cast) id should be Integer")
		return
	}

	cast, err := repositories.cast.FindByID(castID)

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
		c.JSON(http.StatusBadRequest, "file extension is not allowed")
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)

	filePath := fmt.Sprintf("/casts/%d/profile-images", castID)

	file, err := aws.S3.UploadFile(mainImage, filePath, fileExtension)

	if err != nil {
		c.JSON(http.StatusBadRequest, "profileImage should be File")
		return
	}

	cast.ProfileImage = &repository.File{
		Path: *file.Key,
		Size: fileHeader.Size,
	}

	err = repositories.file.Save(cast.ProfileImage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "database error: file save failed")
		log.Fatal(err)
		return
	}

	err = repositories.cast.Update(cast)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "database error: cast update failed")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profileImage": cast.ProfileImage.FullPath(),
	})
}

func Update(c *gin.Context) {
	id, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(cast) id should be Integer")
		return
	}

	var cast Cast
	err = c.BindJSON(&cast)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid JSON")
		return
	}

	castToUpdate, err := repositories.cast.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "cast not found")
		return
	}

	castToUpdate.PersonID = cast.PersonID
	castToUpdate.CharacterID = cast.CharacterID

	err = repositories.cast.Update(castToUpdate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "database error: cast update failed")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cast": castToUpdate,
	})
}
