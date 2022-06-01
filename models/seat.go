package models

type Seat interface {
}

type SeatRepository interface {
	GetSeatsByAreaId(areaId uint) []Seat
}
