package repository

import (
	"gorm.io/gorm"
	"time"
)

type Seat struct {
	UUID              string           `json:"uuid" gorm:"primarykey"`
	AreaBoilerplate   *AreaBoilerplate `json:"areaBoilerplate" gorm:"foreignKey:AreaBoilerplateID"`
	AreaBoilerplateID *uint            `json:"-"`
	Performance       *Performance     `json:"performance" gorm:"foreignKey:PerformanceID"`
	PerformanceID     uint             `json:"-"`
	Grade             *SeatGrade       `json:"seatGrade" gorm:"foreignkey:SeatGradeID"`
	SeatGradeID       uint             `json:"-"`
	Bookings          []SeatBooking    `json:"booking" gorm:"foreignKey:SeatUUID"`
	CreatedAt         time.Time        `json:"createdAt"`
	UpdatedAt         time.Time        `json:"updatedAt"`
	DeletedAt         *gorm.DeletedAt  `json:"-" gorm:"index"`
}

type SeatRepository struct {
	DB *gorm.DB
}

func (s *SeatRepository) GetByAreaAndPerformanceID(areaID uint, performanceID uint) []Seat {
	var seats []Seat

	err := s.DB.
		Joins("Grade").
		Preload("AreaBoilerplate", "area_id = ?", areaID).
		Preload("AreaBoilerplate.Point").
		Where("performance_id = ?", performanceID).
		Order("area_boilerplate_id").
		Find(&seats).
		Error

	if err != nil {
		return nil
	}

	return seats
}

func (s *SeatRepository) GetByScheduleUUID(scheduleUUID string) []Seat {
	var seats []Seat
	err := s.DB.Preload("Grade").Where("schedule_uuid = ?", scheduleUUID).Find(&seats).Error

	if err != nil {
		return nil
	}

	return nil
}

func (s *SeatRepository) GetByUUID(uuid string) *Seat {
	var seat Seat
	err := s.DB.Preload("Grade").Where("uuid = ?", uuid).Find(&seat).Error

	if err != nil {
		return nil
	}

	return &seat

}

func (s *SeatRepository) GetSeatsByAreaIdAndScheduleUUID(areaId uint, scheduleUUID string) []Seat {
	var seats []Seat

	err := s.DB.
		Joins("Grade").
		Preload("Bookings", "schedule_uuid = ?", scheduleUUID).
		Where("area_id = ?", areaId).
		Find(&seats).
		Error

	if err != nil {
		return nil
	}

	return seats
}

func (s *SeatRepository) SaveSeats(seats []Seat) error {
	return s.DB.
		Omit("area_boilerplate_id").
		Save(&seats).Error
}

func (s *SeatRepository) CreateSeats(seats []Seat) error {
	return s.DB.Debug().Create(&seats).Error
}
