package repository

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CastTestSuite struct {
	suite.Suite
	cast           Cast
	castRepository CastRepository
}

func (suite *CastTestSuite) SetupTest() {
	suite.cast = Cast{Person: &Person{ID: 1}, Character: &Character{ID: 1}}
	suite.castRepository = CastRepository{DB: database.Gorm}
}

func (suite *CastTestSuite) TestGetCastsByPerformanceID() {
	casts, err := suite.castRepository.GetCastsByPerformanceID(11)

	if err != nil {
		return
	}

	for _, cast := range casts {
		fmt.Println("id:", cast.ID, "person:", cast.Person, "character:", cast.Character)
	}
}

func TestCastTestSuite(t *testing.T) {
	suite.Run(t, new(CastTestSuite))
}
