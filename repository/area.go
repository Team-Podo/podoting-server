package repository

import (
	"github.com/kwanok/podonine/database"
	"github.com/kwanok/podonine/models"
	"gorm.io/gorm"
)

type Area struct {
	Model
	PlaceId uint `json:"-"`
	Title   string
	Seats   *[]Seat `gorm:"foreignkey:AreaId"`
}

func (area *Area) GetId() uint {
	return area.ID
}

func (area *Area) GetTitle() string {
	return area.Title
}

func (area *Area) GetSeats() []models.Seat {
	var _seats []Seat
	database.Gorm.Model(Seat{AreaId: area.GetId()}).Find(&_seats)

	var seats []models.Seat

	for _, seat := range _seats {
		seats = append(seats, seat)
	}

	return seats
}

type AreaRepository struct {
	Db *gorm.DB
}

func (repo *AreaRepository) Get() []models.Area {
	var _areas []Area
	repo.Db.Model(Area{}).Find(&_areas)

	var areas []models.Area

	for _, area := range _areas {
		areas = append(areas, &area)
	}

	return areas
}

func (repo *AreaRepository) Find(id uint) models.Area {
	var area Area
	result := repo.Db.First(&area, id)
	if result.Error != nil {
		return nil
	}

	return &area
}

func (repo *AreaRepository) Save(area models.Area) models.Area {
	result := repo.Db.Create(area)
	if result.Error != nil {
		return nil
	}

	return area
}
