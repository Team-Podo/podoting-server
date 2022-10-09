package models

type SeatBookingRepository interface {
	Book(uis string, scheduleUUID string, seatUUIDs []string) error
}
