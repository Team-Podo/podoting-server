package repository

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SaveSeatTestSuite struct {
	suite.Suite
	seatRepository *SeatRepository
	seat           Seat
}

type GetSeatTestSuite struct {
	suite.Suite
	seatRepository *SeatRepository
}

func (suite *SaveSeatTestSuite) SetupTest() {
	suite.seatRepository = &SeatRepository{DB: database.Gorm}
	newUUID, _ := uuid.NewUUID()
	suite.seat = Seat{
		UUID:  newUUID.String(),
		Grade: &SeatGrade{ID: 1},
	}
}

func (suite *GetSeatTestSuite) TestGetSeatsByAreaIdAndScheduleUUID() {
	seatRepository := &SeatRepository{DB: database.Gorm}
	seats := seatRepository.GetSeatsByAreaIdAndScheduleUUID(26, "4c4acf6d-460a-11ed-ab11-0a58a9feac02")
	for _, seat := range seats {
		fmt.Println(seat)
	}
}

func TestSaveSeatTestSuite(t *testing.T) {
	suite.Run(t, new(SaveSeatTestSuite))
}

func TestGetSeatTestSuite(t *testing.T) {
	suite.Run(t, new(GetSeatTestSuite))
}
