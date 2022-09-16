package models

import "github.com/Team-Podo/podoting-server/repository"

type SeatRepository interface {
	GetSeatsByAreaIdAndScheduleUUID(areaId uint, scheduleUUID string) []repository.Seat
}
