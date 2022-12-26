package repository

import (
	"database/sql"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type PerformanceTestSuite struct {
	suite.Suite
	performance           Performance
	performanceRepository PerformanceRepository
}

func (suite *PerformanceTestSuite) SetupTest() {
	db, err := gorm.Open(
		sqlite.Open(
			utils.RootPath()+"sqlite/test.db",
		),
		&gorm.Config{},
	)
	suite.NoError(err)

	suite.performanceRepository = PerformanceRepository{DB: db.Debug().Begin()}
}

func (suite *PerformanceTestSuite) makePerformance() Performance {
	performance := Performance{
		Product: &Product{
			Title: "Test Product",
		},
		Thumbnail: &File{
			Size: 0,
			Path: "test thumbnail",
		},
		Place: &Place{
			Name: "Test Place",
			Location: &Location{
				Name:      "Test Location",
				Longitude: 0,
				Latitude:  0,
			},
			PlaceImage: &File{
				Size: 0,
				Path: "",
			},
		},
		Schedules: []Schedule{
			{
				Memo: "",
				Date: "",
				Open: false,
				Casts: []Cast{
					{
						Character: &Character{
							Name: "test character",
						},
						Person: &Person{
							Name: "test person",
						},
						Schedules: nil,
					},
				},
				Time: sql.NullString{
					String: "test time",
					Valid:  false,
				},
			},
		},
		Title:       "Test Performance",
		RunningTime: "240분",
		StartDate:   "2022-10-20",
		EndDate:     "2022-12-31",
		Rating:      "전체관람가",
	}

	err := suite.performanceRepository.Save(&performance)
	suite.NoError(err)

	return performance
}

func (suite *PerformanceTestSuite) makePerformanceByTitle(title string) Performance {
	performance := Performance{
		Product: &Product{
			Title: "Test Product",
		},
		Thumbnail: &File{
			Size: 0,
			Path: "test thumbnail",
		},
		Place: &Place{
			Name: "Test Place",
			Location: &Location{
				Name:      "Test Location",
				Longitude: 0,
				Latitude:  0,
			},
			PlaceImage: &File{
				Size: 0,
				Path: "",
			},
		},
		Schedules: []Schedule{
			{
				Memo: "",
				Date: "",
				Open: false,
				Casts: []Cast{
					{
						Character: &Character{
							Name: "test character",
						},
						Person: &Person{
							Name: "test person",
						},
						Schedules: nil,
					},
				},
				Time: sql.NullString{
					String: "test time",
					Valid:  false,
				},
			},
		},
		Title:       title,
		RunningTime: "240분",
		StartDate:   "2022-10-20",
		EndDate:     "2022-12-31",
		Rating:      "전체관람가",
	}

	err := suite.performanceRepository.Save(&performance)
	suite.NoError(err)

	return performance
}

func (suite *PerformanceTestSuite) TestGetWith() {
	suite.makePerformance()
	suite.makePerformance()

	performances := suite.performanceRepository.GetWith()

	suite.Equal(2, len(performances))
}

func (suite *PerformanceTestSuite) TestGetWithPlace() {
	suite.makePerformance()
	suite.makePerformance()

	performancesPlaceNotExists := suite.performanceRepository.GetWith()

	for _, performance := range performancesPlaceNotExists {
		suite.Nil(performance.Place)
	}

	performancesPlaceExists := suite.performanceRepository.GetWith("Place")

	for _, performance := range performancesPlaceExists {
		suite.NotNil(performance.Place)
	}
}

func (suite *PerformanceTestSuite) TestGetWithForMainPage() {
	suite.makePerformance()
	suite.makePerformance()

	performances := suite.performanceRepository.GetWith(
		"Place",
		"Thumbnail",
		"Place.Location",
	)

	for _, performance := range performances {
		suite.NotNil(performance.Place)
		suite.NotNil(performance.Place.PlaceImage)
		suite.NotNil(performance.Thumbnail)
	}
}

func (suite *PerformanceTestSuite) TestGetWithKeywordForMainPage() {
	suite.makePerformance()
	suite.makePerformance()
	suite.makePerformanceByTitle("안녕하세요")

	performances := suite.performanceRepository.SetKeyword("Test").GetWith(
		"Thumbnail",
		"Place.Location",
	)

	suite.Equal(2, len(performances))

	for _, performance := range performances {
		suite.NotNil(performance.Place)
		suite.NotNil(performance.Thumbnail)
		suite.NotNil(performance.Place.Location)
	}

	performances = suite.performanceRepository.SetKeyword("안녕").GetWith(
		"Thumbnail",
		"Place.Location",
	)

	suite.Equal(1, len(performances))

	for _, performance := range performances {
		suite.NotNil(performance.Place)
		suite.NotNil(performance.Thumbnail)
		suite.NotNil(performance.Place.Location)
	}
}

func TestPerformanceTestSuite(t *testing.T) {
	suite.Run(t, new(PerformanceTestSuite))
}
