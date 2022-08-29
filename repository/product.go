package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"os"
	"strconv"
	"time"
)

type Product struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	Title     string          `json:"title"`
	FileID    uint            `json:"-"`
	File      *File           `json:"file" gorm:"foreignkey:FileID"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}

func (product *Product) IsNil() bool {
	if product == nil {
		return true
	}

	return false
}

func (product *Product) IsNotNil() bool {
	if product == nil {
		return false
	}

	return true
}

func (product *Product) GetFileURL() string {
	if product.File.IsNil() {
		return ""
	}

	return os.Getenv("CDN_URL") + "/" + product.File.Path
}

type ProductRepository struct {
	DB *gorm.DB
}

func (repo *ProductRepository) GetWithQueryMap(query map[string]any) ([]Product, error) {
	db := repo.DB

	if query["reversed"] == true {
		db = db.Order("id desc")
	}

	if query["limit"] != nil {
		limit, _ := strconv.Atoi(utils.ToString(query["limit"]))
		db = db.Limit(limit)
	}

	if query["offset"] != nil {
		offset, _ := strconv.Atoi(utils.ToString(query["offset"]))
		db = db.Offset(offset)
	}

	var products []Product

	err := db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepository) FindByID(id uint) (*Product, error) {
	var product Product
	product.ID = id
	err := repo.DB.Preload("File").First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) Save(product *Product) error {
	err := repo.DB.Create(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) Update(product *Product) error {
	if product.File != nil {
		product.FileID = product.File.ID
	}

	err := repo.DB.Model(&Product{ID: product.ID}).Updates(product).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) Delete(id uint) error {
	var product Product
	product.ID = id
	err := repo.DB.Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) GetTotalWithQueryMap(query map[string]any) (int64, error) {
	db := repo.DB

	if query["reversed"] == true {
		db = db.Order("id desc")
	}

	if query["limit"] != nil {
		limit, _ := strconv.Atoi(utils.ToString(query["limit"]))
		db = db.Limit(limit)
	}

	if query["offset"] != nil {
		offset, _ := strconv.Atoi(utils.ToString(query["offset"]))
		db = db.Offset(offset)
	}

	var count int64
	err := db.Model(&Product{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
