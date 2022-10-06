package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AreaBoilerplate struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name"`
	Area      *Area          `json:"area" gorm:"foreignkey:AreaID"`
	AreaID    uint           `json:"-"`
	Seats     []Seat         `json:"seats" gorm:"foreignkey:AreaBoilerplateID"`
	Point     *Point         `json:"point" gorm:"foreignkey:PointID"`
	PointID   uint           `json:"-"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type AreaBoilerplateRepository struct {
	DB *gorm.DB
}

func (r *AreaBoilerplateRepository) GetAreaBoilerplateByAreaID(areaID uint) ([]AreaBoilerplate, error) {
	var areaBoilerplates []AreaBoilerplate

	err := r.DB.Where("area_id = ?", areaID).Find(&areaBoilerplates).Error
	fmt.Println(areaBoilerplates)

	return areaBoilerplates, err
}

func (r *AreaBoilerplateRepository) SaveSeats(boilerplates []AreaBoilerplate, performanceID uint) error {
	var seats []Seat
	for _, boilerplate := range boilerplates {
		bpID := boilerplate.ID
		seats = append(seats, Seat{
			UUID:              uuid.New().String(),
			AreaBoilerplateID: &bpID,
			PerformanceID:     performanceID,
			SeatGradeID:       1,
		})
	}

	return r.DB.Create(&seats).Error
}
