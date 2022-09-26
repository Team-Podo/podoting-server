package repository

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PlaceTestSuite struct {
	suite.Suite
	place           Place
	placeRepository PlaceRepository
}

func (suite *PlaceTestSuite) SetupTest() {
	suite.place = Place{
		Name:       "서울",
		Location:   nil,
		LocationID: nil,
	}
	suite.placeRepository = PlaceRepository{DB: database.Gorm}
}

func (suite *PlaceTestSuite) TestCreate() {
	_ = suite.placeRepository.Create(&suite.place)

	suite.T().Log(suite.place.ID)
}

func TestPlaceTestSuite(t *testing.T) {
	suite.Run(t, new(PlaceTestSuite))
}
