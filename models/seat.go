package models

import "github.com/Team-Podo/podoting-server/repository"

type SeatRepository interface {
	GetByAreaID(areaID uint) []repository.Seat
	GetSeatsByAreaIdAndScheduleUUID(areaId uint, scheduleUUID string) []repository.Seat
	SaveSeats(seats []repository.Seat) error
}
