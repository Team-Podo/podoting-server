package models

import "github.com/Team-Podo/podoting-server/repository"

type AreaRepository interface {
	GetByPlaceID(placeID uint) []repository.Area
	FindOne(placeID uint, areaID uint) *repository.Area
	Create(area *repository.Area) error
	Update(area *repository.Area) error
	Delete(area *repository.Area) error
	GetBackgroundImageByAreaId(areaID uint) string
}
