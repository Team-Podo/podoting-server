package repository

import (
	"github.com/Team-Podo/podoting-server/models"
	"gorm.io/gorm"
)

type Seat struct {
	Model
	AreaId uint `json:"-"`
	Code   string
}

type SeatRepository struct {
	Db *gorm.DB
}

func (s *Seat) GetId() uint {
	return s.ID
}

func (repo *SeatRepository) GetSeatsByAreaId(areaId uint) []models.Seat {
	var _seats []Seat
	repo.Db.Model(Seat{AreaId: areaId}).Find(&_seats)

	var seats []models.Seat

	for _, seat := range _seats {
		seats = append(seats, seat)
	}

	return seats
}
