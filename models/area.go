package models

type Area interface {
	GetId() uint
	GetTitle() string
	GetSeats() []Seat
}

type AreaRepository interface {
	Get() []Area
	Find(id uint) Area
	Save(area Area) Area
}
