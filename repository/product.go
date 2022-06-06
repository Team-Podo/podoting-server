package repository

import (
	"fmt"
	"github.com/kwanok/podonine/models"
	"gorm.io/gorm"
)

type Product struct {
	Model
	Title   string
	Place   Place `gorm:"foreignkey:PlaceId"`
	PlaceId uint  `json:"-"`
}

func (product *Product) GetId() uint {
	return product.ID
}

func (product *Product) GetTitle() string {
	return product.Title
}

type ProductRepository struct {
	Db *gorm.DB
}

func (repo *ProductRepository) GetProducts() []models.Product {
	var _products []Product
	repo.Db.Model(Product{}).Find(&_products)

	var products []models.Product

	for _, product := range _products {
		products = append(products, &product)
	}

	return products
}

func (repo *ProductRepository) GetProductById(id uint) models.Product {
	var product Product
	product.ID = id
	result := repo.Db.Preload("Place.Areas.Seats").First(&product)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil
	}

	return &product
}

func (repo *ProductRepository) SaveProduct(product models.Product) models.Product {
	var _product Product
	_product.Title = product.GetTitle()

	result := repo.Db.Create(&_product)
	if result.Error != nil {
		return nil
	}

	return &_product
}

func (repo *ProductRepository) DeleteProductById(id uint) {
	var product Product
	product.ID = id
	repo.Db.Delete(&product)
}
