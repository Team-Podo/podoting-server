package models

import "github.com/Team-Podo/podoting-server/repository"

type SeatBookingRepository interface {
	Get(userUID string, scheduleUUID string, seatUUIDs []string) ([]repository.SeatBooking, error)
	Book(userUID string, scheduleUUID string, seatUUIDs []string) error
}
