package repository

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
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
	suite.productRepository = &ProductRepository{DB: database.Gorm}
	suite.product = Product{Title: "수정 전"}
}

func (suite *ProductTestSuite) TestGet() {
	products, _ := suite.productRepository.GetWithQueryMap(map[string]any{})
	for _, product := range products {
		fmt.Println("id:", product.ID, "title", product.Title)
	}

	assert.NotNil(suite.T(), products)
}

func (suite *ProductTestSuite) TestGetProductById() {
	product, _ := suite.productRepository.FindByID(1)

	assert.Equal(suite.T(), uint(1), product.ID)
}

func (suite *ProductTestSuite) TestSaveProduct() {
	err := suite.productRepository.Save(&suite.product)
	assert.Nil(suite.T(), err)

	_product, _ := suite.productRepository.FindByID(suite.product.ID)

	assert.Equal(suite.T(), "테스트 상품", _product.Title)
}

func (suite *ProductTestSuite) TestUpdateProduct() {
	err := suite.productRepository.Save(&suite.product)
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), "수정 전", suite.product.Title)
	fmt.Println(suite.product)

	var _product Product
	_product.ID = suite.product.ID
	_product.Title = "수정 후"

	err = suite.productRepository.Update(&_product)
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), "수정 후", suite.product.Title)
	fmt.Println(suite.product)
}

func (suite *ProductTestSuite) TestDeleteProductById() {
	err := suite.productRepository.Save(&suite.product)
	assert.Nil(suite.T(), err)

	err = suite.productRepository.Delete(suite.product.ID)
	assert.Nil(suite.T(), err)

	_product, _ := suite.productRepository.FindByID(suite.product.ID)

	assert.Equal(suite.T(), nil, _product)
}

func TestSaveProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
