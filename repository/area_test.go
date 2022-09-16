package repository

import (
	"encoding/json"
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SaveAreaTestSuite struct {
	suite.Suite
	areaRepository *AreaRepository
	area           Area
}

type GetAreaTestSuite struct {
	suite.Suite
	areaRepository *AreaRepository
}

func (suite *SaveAreaTestSuite) SetupTest() {
	suite.areaRepository = &AreaRepository{DB: database.Gorm}
	suite.area = Area{
		Name:    "test",
		PlaceID: 8,
	}
}

func (suite *GetAreaTestSuite) SetupTest() {
	suite.areaRepository = &AreaRepository{DB: database.Gorm}
}

func (suite *GetAreaTestSuite) TestFindOneArea() {
	area := suite.areaRepository.FindOne(1)
	jsonArea, _ := json.Marshal(&area)
	fmt.Println(string(jsonArea))
	fmt.Println(area)
	suite.NotNil(area)
}

func (suite *SaveAreaTestSuite) TestSaveArea() {
	err := suite.areaRepository.SaveArea(&suite.area)
	suite.NotNil(err)
}

func TestSaveAreaTestSuite(t *testing.T) {
	suite.Run(t, new(SaveAreaTestSuite))
}

func TestGetAreaTestSuite(t *testing.T) {
	suite.Run(t, new(GetAreaTestSuite))
}
