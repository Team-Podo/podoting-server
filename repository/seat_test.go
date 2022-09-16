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
	suite.seatRepository = &SeatRepository{Db: database.Gorm}
	newUUID, _ := uuid.NewUUID()
	suite.seat = Seat{
		UUID:   newUUID.String(),
		AreaID: 1,
		Grade:  &SeatGrade{ID: 1},
		Point: &Point{
			X: float64(1),
			Y: float64(1),
		},
	}
}

func (suite *SaveSeatTestSuite) TestSaveSeats() {
	fmt.Println(suite.seat)
	err := suite.seatRepository.SaveSeats([]Seat{suite.seat})
	suite.Nil(err)
}

func (suite *GetSeatTestSuite) TestGetSeatsByAreaIdAndScheduleUUID() {
	seatRepository := &SeatRepository{Db: database.Gorm}
	seats := seatRepository.GetSeatsByAreaIdAndScheduleUUID(4, "18a08d6d-5d71-4d5b-aff3-d7f421312ee4")
	fmt.Println(seats)
	for i, seat := range seats {
		fmt.Println(i, seat, seat.Grade, seat.Point, seat.Bookings)
	}
}

func TestSaveSeatTestSuite(t *testing.T) {
	suite.Run(t, new(SaveSeatTestSuite))
}

func TestGetSeatTestSuite(t *testing.T) {
	suite.Run(t, new(GetSeatTestSuite))
}
