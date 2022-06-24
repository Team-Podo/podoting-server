package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductTestSuite struct {
	suite.Suite
	productRepository *ProductRepository
	product           Product
}

func (suite *ProductTestSuite) SetupTest() {
	suite.productRepository = &ProductRepository{Db: Gorm}
	suite.product = Product{Title: "수정 전"}
}

func (suite *ProductTestSuite) TestGet() {
	products := suite.productRepository.Get()
	for _, product := range products {
		fmt.Println(product)
	}

	assert.NotNil(suite.T(), products)
}

func (suite *ProductTestSuite) TestGetProductById() {
	product := suite.productRepository.GetProductById(1)
	fmt.Println(product.GetId())
	fmt.Println(product.GetTitle())
	fmt.Println(product.GetPlace())

	assert.Equal(suite.T(), uint(1), product.GetId())
}

func (suite *ProductTestSuite) TestSaveProduct() {
	product := suite.productRepository.SaveProduct(&Product{
		Title: "테스트 상품",
	})

	_product := suite.productRepository.GetProductById(product.GetId())

	assert.Equal(suite.T(), "테스트 상품", _product.GetTitle())
}

func (suite *ProductTestSuite) TestUpdateProduct() {
	product := suite.productRepository.SaveProduct(&suite.product)
	assert.Equal(suite.T(), "수정 전", product.GetTitle())
	fmt.Println(product)

	var _product Product
	_product.ID = product.GetId()
	_product.Title = "수정 후"

	product = suite.productRepository.Update(&_product)
	assert.Equal(suite.T(), "수정 후", product.GetTitle())
	fmt.Println(product)
}

func (suite *ProductTestSuite) TestDeleteProductById() {
	product := suite.productRepository.SaveProduct(&Product{
		Title: "테스트 삭제 상품",
	})

	suite.productRepository.DeleteProductById(product.GetId())

	_product := suite.productRepository.GetProductById(product.GetId())

	assert.Equal(suite.T(), nil, _product)
}

func TestSaveProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
