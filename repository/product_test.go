package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductTestSuite struct {
	suite.Suite
	productRepository *ProductRepository
}

func (suite *ProductTestSuite) SetupTest() {
	suite.productRepository = &ProductRepository{Db: Gorm}
}

func (suite *ProductTestSuite) TestGet() {
	products := suite.productRepository.Get()

	assert.NotNil(suite.T(), products)
}

func (suite *ProductTestSuite) TestGetProductById() {
	product := suite.productRepository.GetProductById(1)

	assert.Equal(suite.T(), uint(1), product.GetId())
}

func (suite *ProductTestSuite) TestSaveProduct() {
	product := suite.productRepository.SaveProduct(&Product{
		Title: "테스트 상품",
	})

	_product := suite.productRepository.GetProductById(product.GetId())

	assert.Equal(suite.T(), "테스트 상품", _product.GetTitle())
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
