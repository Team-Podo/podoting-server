package endpoints

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProductTestSuite struct {
	suite.Suite
	product           models.Product
	productRepository models.ProductRepository
}

func (suite *ProductTestSuite) SetupTest() {
	var content = []*repository.ProductContent{
		{},
		{},
	}

	suite.product = &repository.Product{
		Title: "Test1",
		Place: &repository.Place{
			Title: "아트센터",
		},
		Content: content,
	}

	suite.productRepository = &repository.ProductRepository{Db: database.Gorm}
}

func (suite *ProductTestSuite) TestSaveProduct() {
	product := suite.productRepository.Save(suite.product)

	assert.Equal(suite.T(), product.GetTitle(), suite.product.GetTitle())
}

func TestSaveProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
