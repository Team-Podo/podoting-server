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
	BackgroundImageID uint           `json:"-"`
	Seats             []Seat         `json:"seats" gorm:"foreignKey:AreaID"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

type AreaRepository struct {
	DB *gorm.DB
}

func (r *AreaRepository) FindOne(id uint) *Area {
	var area Area
	err := r.DB.
		Debug().
		Preload("Seats.Point").
		Preload("Seats.Bookings").
		Preload("Seats.Grade").
		First(&area, id).
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

func (r *AreaRepository) Update(area *Area) error {
	return r.DB.Save(area).Error
}

func (r *AreaRepository) SaveArea(area *Area) interface{} {
	err := r.DB.Debug().Create(area).Error

	if err != nil {
		return nil
	}

	return area
}
