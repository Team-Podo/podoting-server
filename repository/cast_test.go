package repository

import (
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

func TestCastTestSuite(t *testing.T) {
	suite.Run(t, new(CastTestSuite))
}
