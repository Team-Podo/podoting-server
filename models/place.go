package models

import "github.com/Team-Podo/podoting-server/repository"

type Place interface {
	GetId() uint
	GetTitle() string
}

type PlaceRepository interface {
	FindByID(id uint) (*repository.Place, error)
	Update(place *repository.Place) error
}
