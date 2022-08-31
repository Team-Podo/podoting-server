package repository

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PerformanceContentTestSuite struct {
	suite.Suite
	performance                  *Performance
	performanceRepository        *PerformanceRepository
	PerformanceContentRepository *PerformanceContentRepository
}

func (suite *PerformanceContentTestSuite) SetupTest() {
	performanceContents := []*PerformanceContent{
		{UUID: uuid.New().String(), Title: "이용 안내", Content: "<p>안녕하세요!</p>", Priority: 1},
		{UUID: uuid.New().String(), Title: "공연 안내", Content: "<p>잘 보고 가세요!</p>", Priority: 2},
	}

	suite.performance = &Performance{
		Title:     "TEST 퍼포먼스",
		StartDate: "2022-07-19",
		EndDate:   "2022-08-20",
		Product:   &Product{ID: 1},
		Contents:  performanceContents,
		Place: &Place{
			Name:       "seoul",
			PlaceImage: nil,
			Location: &Location{
				Name:      "서울",
				Longitude: 127.0,
				Latitude:  37.0,
			},
		},
	}
	suite.performanceRepository = &PerformanceRepository{DB: database.Gorm}
	suite.PerformanceContentRepository = &PerformanceContentRepository{DB: database.Gorm}
}

func (suite *PerformanceContentTestSuite) TestSave() {
	_ = suite.performanceRepository.Save(suite.performance)

	for _, content := range suite.performance.Contents {
		fmt.Println("id:", content.UUID, "title:", content.Title)
	}
}

func (suite *PerformanceContentTestSuite) TestUpdate() {
	performance := suite.performanceRepository.FindByID(11)

	_ = suite.PerformanceContentRepository.Save(&PerformanceContent{PerformanceID: performance.ID, UUID: uuid.New().String(), Title: "이용 안내", Content: "<p>안녕하세요!</p>", Priority: 1})
	_ = suite.PerformanceContentRepository.Save(&PerformanceContent{PerformanceID: performance.ID, UUID: uuid.New().String(), Title: "공연 안내", Content: "<p>잘 보다 가세요~</p>", Priority: 2})

	for _, content := range performance.Contents {
		fmt.Println("id:", content.UUID, "title:", content.Title)
	}
}

func TestPerformanceContentTestSuite(t *testing.T) {
	suite.Run(t, new(PerformanceContentTestSuite))
}
