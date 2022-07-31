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
	suite.productRepository = &ProductRepository{Db: database.Gorm}
	suite.product = Product{Title: "수정 전"}
}

func (suite *ProductTestSuite) TestGet() {
	products := suite.productRepository.Get(map[string]any{})
	for _, product := range products {
		fmt.Println("id:", product.GetId(), "title", product.GetTitle())
	}

	assert.NotNil(suite.T(), products)
}

func (suite *ProductTestSuite) TestGetProductById() {
	product := suite.productRepository.Find(1)
	fmt.Println(product.GetId())
	fmt.Println(product.GetTitle())
	fmt.Println(product.GetPlace())

	assert.Equal(suite.T(), uint(1), product.GetId())
}

func (suite *ProductTestSuite) TestSaveProduct() {
	product := suite.productRepository.Save(&Product{
		Title: "테스트 상품",
	})

	_product := suite.productRepository.Find(product.GetId())

	assert.Equal(suite.T(), "테스트 상품", _product.GetTitle())
}

func (suite *ProductTestSuite) TestUpdateProduct() {
	product := suite.productRepository.Save(&suite.product)
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
	product := suite.productRepository.Save(&Product{
		Title: "테스트 삭제 상품",
	})

	suite.productRepository.Delete(product.GetId())

	_product := suite.productRepository.Find(product.GetId())

	assert.Equal(suite.T(), nil, _product)
}

func TestSaveProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
