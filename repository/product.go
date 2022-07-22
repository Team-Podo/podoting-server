package repository

import (
	"fmt"
	"github.com/kwanok/podonine/models"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"strconv"
)

type Product struct {
	Model
	Title   string
	Place   *Place `gorm:"foreignkey:PlaceId"`
	PlaceId uint   `json:"-"`
	Content []*ProductContent
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

func (repo *ProductRepository) Update(product models.Product) models.Product {
	var _product Product
	_product.ID = product.GetId()
	repo.Db.First(&_product)

	_product.Title = product.GetTitle()
	repo.Db.Save(&_product)

	return product
}

func (repo *ProductRepository) Delete(id uint) {
	var product Product
	product.ID = id
	repo.Db.Delete(&product)
}
