package models

import "github.com/Team-Podo/podoting-server/repository"

type PlaceRepository interface {
	FindByID(id uint) (*repository.Place, error)
	FindAll() ([]repository.Place, error)
	Create(place *repository.Place) error
	Update(place *repository.Place) error
	Delete(placeID uint) error
}
