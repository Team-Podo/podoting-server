package models

import "github.com/Team-Podo/podoting-server/repository"

type Area interface {
	GetId() uint
	GetTitle() string
	GetSeats() []repository.Seat
}

type AreaRepository interface {
	FindOne(id uint) *repository.Area
	SaveArea(area *repository.Area) interface{}
	Update(area *repository.Area) error
	GetBackgroundImageByAreaId(areaID uint) string
}
