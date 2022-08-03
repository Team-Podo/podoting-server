package product

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/Team-Podo/podoting-server/utils/aws"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Product struct {
	ID    uint
	Title string `json:"title"`
}

var repositories Repository

type Repository struct {
	product models.ProductRepository
	file    models.FileRepository
}

func init() {
	repositories = Repository{
		product: &repository.ProductRepository{DB: database.Gorm},
		file:    &repository.FileRepository{DB: database.Gorm},
	}
}

func Get(c *gin.Context) {
	// ------ 쿼리스트링 검증 Start ------

	limitQuery := c.Query("limit")
	offsetQuery := c.Query("offset")
	reversedQuery := c.Query("reversed")

	var limit int
	var offset int
	var reversed = false
	var err error

	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, "limit should be Integer")
			return
		}
	}

	if offsetQuery != "" {
		offset, err = strconv.Atoi(offsetQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, "offset should be Integer")
			return
		}
	}

	if reversedQuery != "" {
		reversed = true
	}

	query := map[string]any{
		"limit":    limit,
		"offset":   offset,
		"reversed": reversed,
	}

	// ------ 쿼리스트링 검증 End ------

	// ------ 상품 가져오기 Start ------

	products, err := repositories.product.GetWithQueryMap(query)

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	// ------ 상품 가져오기 End ------

	// ------ 응답 폼 만들기 Start ------

	var productResponses []utils.MapSlice

	for _, product := range products {
		productResponses = append(productResponses, utils.MapSlice{
			utils.MapItem{Key: "id", Value: product.ID},
			utils.MapItem{Key: "title", Value: product.Title},
			utils.MapItem{Key: "createdAt", Value: product.CreatedAt},
			utils.MapItem{Key: "updatedAt", Value: product.UpdatedAt},
		})
	}

	total, err := repositories.product.GetTotalWithQueryMap(query)

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	// ------ 응답 폼 만들기 End ------

	c.JSON(http.StatusOK, map[string]any{
		"products": productResponses,
		"total":    total,
	})
}

func Find(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	product, err := repositories.product.FindByID(uint(intId))

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, utils.MapSlice{
		utils.MapItem{Key: "id", Value: product.ID},
		utils.MapItem{Key: "title", Value: product.Title},
		utils.MapItem{Key: "file", Value: product.File},
		utils.MapItem{Key: "createdAt", Value: product.CreatedAt},
		utils.MapItem{Key: "updatedAt", Value: product.UpdatedAt},
	})
}

func Create(c *gin.Context) {
	var json Product
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	product := repository.Product{
		Title: json.Title,
	}

	err := repositories.product.Save(&product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, product)
}

func UploadMainImage(c *gin.Context) {
	id, err := getIntParam(c, "id")

	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	product, err := repositories.product.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}

	mainImage, fileHeader, err := c.Request.FormFile("mainImage")

	if err != nil {
		c.JSON(http.StatusBadRequest, "mainImage is required")
		return
	}

	if !checkFileExtension(fileHeader) {
		c.JSON(http.StatusBadRequest, "File extension is not allowed")
		return
	}

	fileExtension := filepath.Ext(fileHeader.Filename)

	file, err := aws.S3.UploadFile(mainImage, "/main-images", fileExtension)

	if err != nil {
		c.JSON(http.StatusBadRequest, "mainImage should be File")
		return
	}

	product.File = &repository.File{
		Path: *file.Key,
		Size: fileHeader.Size,
	}

	err = repositories.file.Save(product.File)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	err = repositories.product.Update(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"mainImage": getFullURLFromFile(product.File),
	})
}

func checkFileExtension(fileHeader *multipart.FileHeader) bool {
	extension := strings.ToLower(filepath.Ext(fileHeader.Filename))
	return extension == ".jpg" || extension == ".jpeg" || extension == ".png"
}

func getIntParam(c *gin.Context, param string) (int, error) {
	id := c.Param(param)

	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return intId, nil
}

func getFullURLFromFile(file *repository.File) string {
	return os.Getenv("CDN_URL") + "/" + file.Path
}
