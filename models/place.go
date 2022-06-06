package models

type Place interface {
	GetId() uint
	GetTitle() string
}

type PlaceRepository interface {
	Get() []Place
	Find(id uint) Place
	Save(place Place) Place
}
