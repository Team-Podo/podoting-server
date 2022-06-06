package repository

import (
	"fmt"
	"github.com/kwanok/podonine/models"
	"gorm.io/gorm"
)

type Place struct {
	Model
	Title    string
	Products *[]Product `gorm:"foreignkey:PlaceId"`
	Areas    *[]Area    `gorm:"foreignkey:PlaceId"`
}

func (place *Place) GetId() uint {
	return place.ID
}

func (place *Place) GetTitle() string {
	return place.Title
}

type PlaceRepository struct {
	Db *gorm.DB
}

func (repo *PlaceRepository) Get() []models.Place {
	var _places []Place
	repo.Db.Model(Place{}).Find(&_places)

	var places []models.Place

	for _, place := range _places {
		places = append(places, &place)
	}

	return places
}

func (repo *PlaceRepository) Find(id uint) models.Place {
	var place Place
	place.ID = id
	result := repo.Db.Preload("Areas.Seats").First(&place)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil
	}

	return &place
}

func (repo *PlaceRepository) Save(place models.Place) models.Place {
	var _place Place
	_place.Title = place.GetTitle()

	result := repo.Db.Create(&_place)
	if result.Error != nil {
		return nil
	}

	return &_place
}
