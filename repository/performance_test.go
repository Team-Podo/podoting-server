package repository

import (
	"fmt"
	"github.com/kwanok/podonine/database"
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
	suite.performance = Performance{Title: "2022 서울투어", StartDate: "2022-07-19", EndDate: "2022-08-20"}
	suite.performanceRepository = PerformanceRepository{Db: database.Gorm}
}

func (suite *PerformanceTestSuite) TestGet() {
	products := suite.performanceRepository.Get(map[string]any{})
	for _, product := range products {
		fmt.Println("id:", product.GetId(), "title:", product.GetTitle())
	}
}

func (suite *PerformanceTestSuite) TestFind() {
	product := suite.performanceRepository.Find(1)
	fmt.Println("id:", product.GetId(), "title:", product.GetTitle())

	assert.Equal(suite.T(), uint(1), product.GetId())
}

func (suite *PerformanceTestSuite) TestFindNotFound() {
	product := suite.performanceRepository.Find(10)
	fmt.Println("product", product)

	assert.Equal(suite.T(), nil, product)
}

func (suite *PerformanceTestSuite) TestSave() {
	result := suite.performanceRepository.Save(&suite.performance)

	fmt.Println(result.GetId())
}

func (suite *PerformanceTestSuite) TestUpdate() {
	suite.performance.ID = 2
	suite.performance.Title = "2022 서울투어 고고~"
	result := suite.performanceRepository.Update(&suite.performance)

	fmt.Println(result.GetId())
}

func (suite *PerformanceTestSuite) TestDelete() {
	suite.performanceRepository.Delete(2)
}

func TestPerformanceTestSuite(t *testing.T) {
	suite.Run(t, new(PerformanceTestSuite))
}
