package models

type Area interface {
	GetId() uint
	GetTitle() string
	GetSeats() []Seat
}

type AreaRepository interface {
	GetAreaById(id uint) Area
	SaveArea(area Area) Area
	DeleteAreaById(id uint)
}
