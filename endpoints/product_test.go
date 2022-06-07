package endpoints

import (
	"github.com/go-playground/assert/v2"
	"github.com/kwanok/podonine/models"
	"github.com/kwanok/podonine/repository"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductTestSuite struct {
	suite.Suite
	product           models.Product
	productRepository models.ProductRepository
}

func (suite *ProductTestSuite) SetupTest() {
	suite.product = &repository.Product{
		Title: "Test1",
		Place: repository.Place{},
	}

	suite.productRepository = &repository.ProductRepository{Db: repository.Gorm}
}

func (suite *ProductTestSuite) TestSaveProduct() {
	product := suite.productRepository.SaveProduct(suite.product)

	assert.Equal(suite.T(), product.GetTitle(), suite.product.GetTitle())
}

func TestSaveProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
