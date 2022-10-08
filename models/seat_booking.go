package models

type SeatBookingRepository interface {
	Book(scheduleUUID string, seatUUIDs []string) error
}
