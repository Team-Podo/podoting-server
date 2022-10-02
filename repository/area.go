package repository

import (
	"gorm.io/gorm"
	"time"
)

type Area struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	Name              string         `json:"name"`
	Place             *Place         `json:"place" gorm:"foreignKey:PlaceID"`
	PlaceID           uint           `json:"-"`
	BackgroundImage   *File          `json:"backgroundImage" gorm:"foreignKey:BackgroundImageID"`
	BackgroundImageID *uint          `json:"-"`
	Seats             []Seat         `json:"seats" gorm:"foreignKey:AreaID"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

type AreaRepository struct {
	DB *gorm.DB
}

func (r *AreaRepository) GetByPlaceID(placeID uint) []Area {
	var areas []Area
	err := r.DB.Joins("BackgroundImage").Where("place_id = ?", placeID).
		Find(&areas).Error

	if err != nil {
		return nil
	}

	return areas
}

func (r *AreaRepository) FindOne(placeID uint, areaID uint) *Area {
	var area Area
	err := r.DB.
		Preload("Seats.Point").
		Preload("Seats.Bookings").
		Preload("Seats.Grade").
		Joins("BackgroundImage").
		Where("place_id = ?", placeID).
		First(&area, areaID).
		Error

	if err != nil {
		return nil
	}

	return &area
}

func (r *AreaRepository) GetBackgroundImageByAreaId(areaID uint) string {
	var area Area
	err := r.DB.
		Joins("BackgroundImage").
		First(&area, areaID).
		Error

	if err != nil {
		return ""
	}

	if area.BackgroundImage.IsNil() {
		return ""
	}

	return area.BackgroundImage.FullPath()
}

func (r *AreaRepository) Create(area *Area) error {
	return r.DB.Create(area).Error
}

func (r *AreaRepository) Update(area *Area) error {
	return r.DB.Save(area).Error
}

func (r *AreaRepository) Delete(area *Area) error {
	return r.DB.Delete(area).Error
}
