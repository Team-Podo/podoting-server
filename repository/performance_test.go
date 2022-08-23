package repository

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PerformanceTestSuite struct {
	suite.Suite
	performance           Performance
	performanceRepository PerformanceRepository
}

func (suite *PerformanceTestSuite) SetupTest() {
	suite.performance = Performance{Title: "2022 서울투어", StartDate: "2022-07-19", EndDate: "2022-08-20", Product: &Product{ID: 1}}
	suite.performanceRepository = PerformanceRepository{DB: database.Gorm}
}

func (suite *PerformanceTestSuite) TestGet() {
	products := suite.performanceRepository.GetWithQueryMap(map[string]any{})
	for _, product := range products {
		fmt.Println("id:", product.ID, "title:", product.Title)
	}
}

func (suite *PerformanceTestSuite) TestFind() {
	product := suite.performanceRepository.FindByID(1)
	fmt.Println("id:", product.ID, "title:", product.Title)

	assert.Equal(suite.T(), uint(1), product.ID)
}

func (suite *PerformanceTestSuite) TestFindNotFound() {
	product := suite.performanceRepository.FindByID(10)
	fmt.Println("product", product)

	assert.Equal(suite.T(), nil, product)
}

func (suite *PerformanceTestSuite) TestSave() {
	_ = suite.performanceRepository.Save(&suite.performance)

	fmt.Println(suite.performance.ID)
}

func (suite *PerformanceTestSuite) TestUpdate() {
	suite.performance.ID = 10
	suite.performance.Title = "2022 서울투어 고고~"
	_ = suite.performanceRepository.Update(&suite.performance)

	fmt.Println(suite.performance.ID)
}

func (suite *PerformanceTestSuite) TestDelete() {
	_ = suite.performanceRepository.Delete(10)
}

func TestPerformanceTestSuite(t *testing.T) {
	suite.Run(t, new(PerformanceTestSuite))
}
