package repository

import (
	"errors"
	"fmt"
	"github.com/Team-Podo/podoting-server/models"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"strconv"
	"time"
)

type Product struct {
	ID        uint `json:"id" gorm:"primarykey"`
	Title     string
	Place     *Place `gorm:"foreignkey:PlaceId"`
	PlaceId   uint   `json:"-"`
	Content   []*ProductContent
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}

func (product *Product) GetId() uint {
	return product.ID
}

func (product *Product) GetTitle() string {
	return product.Title
}

func (product *Product) GetPlace() models.Place {
	return product.Place
}

func (product *Product) GetCreatedAt() string {
	return product.CreatedAt.Format("2006-01-02 15:04:05")
}

func (product *Product) GetUpdatedAt() string {
	return product.UpdatedAt.Format("2006-01-02 15:04:05")
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

type ProductRepository struct {
	Db *gorm.DB
}

func (repo *ProductRepository) Get(query map[string]any) []models.Product {
	var _products []*Product

	db := repo.Db

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

	db.Find(&_products)

	var products = make([]models.Product, len(_products))
	for i, _product := range _products {
		products[i] = _product
	}

	return products
}

func (repo *ProductRepository) Find(id uint) models.Product {
	var product Product
	product.ID = id
	result := repo.Db.Preload("Place.Areas.Seats").First(&product)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil
	}

	return &product
}

func (repo *ProductRepository) Save(product models.Product) models.Product {
	var _product Product
	_product.Title = product.GetTitle()

	result := repo.Db.Create(&_product)
	if result.Error != nil {
		return nil
	}

	return &_product
}

func (repo *ProductRepository) Update(productModel models.Product) models.Product {
	err := repo.Db.First(&Product{}, productModel.GetId()).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	repo.Db.Model(&Product{ID: productModel.GetId()}).Updates(Product{
		Title: productModel.GetTitle(),
	})

	return productModel
}

func (repo *ProductRepository) Delete(id uint) {
	var product Product
	product.ID = id
	repo.Db.Delete(&product)
}

func (repo *ProductRepository) GetTotal(query map[string]any) int64 {
	db := repo.Db

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
	db.Model(&Product{}).Count(&count)

	return count
}
