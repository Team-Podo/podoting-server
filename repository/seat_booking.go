package repository

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type SeatBooking struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Seat         *Seat          `json:"seat" gorm:"foreignKey:SeatUUID"`
	SeatUUID     string         `json:"seatUUID" gorm:"size:36"`
	Schedule     *Schedule      `json:"schedule" gorm:"foreignKey:ScheduleUUID"`
	ScheduleUUID string         `json:"scheduleUUID" gorm:"size:36"`
	BookerUID    string         `json:"bookerUID" gorm:"size:36"`
	Booked       bool           `json:"booked"`
	Canceled     bool           `json:"canceled"`
	BookedAt     time.Time      `json:"bookedAt"`
	CanceledAt   *time.Time     `json:"canceledAt"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type SeatBookingRepository struct {
	DB *gorm.DB
}

func (s *SeatBookingRepository) Get(userUID string, scheduleUUID string, seatUUIDs []string) ([]SeatBooking, error) {
	var seatBookings []SeatBooking
	err := s.DB.
		Preload("Seat.Grade").
		Where("schedule_uuid = ? AND seat_uuid IN ? AND booker_uid = ? AND canceled = false", scheduleUUID, seatUUIDs, userUID).
		Find(&seatBookings).Error

	if err != nil {
		return nil, err
	}

	return seatBookings, nil
}

func (s *SeatBookingRepository) Book(uid string, scheduleUUID string, seatUUIDs []string) error {
	var seatBookings []SeatBooking
	err := s.DB.
		Where("schedule_uuid = ? AND seat_uuid IN ? AND canceled = false", scheduleUUID, seatUUIDs).
		Find(&seatBookings).Error
	if err != nil {
		return err
	}

	if len(seatBookings) > 0 {
		return errors.New("already booked")
	}

	for _, seatUUID := range seatUUIDs {
		seatBooking := SeatBooking{
			ScheduleUUID: scheduleUUID,
			BookerUID:    uid,
			SeatUUID:     seatUUID,
			Booked:       true,
			BookedAt:     time.Now(),
		}
		err = s.DB.Create(&seatBooking).Error
		if err != nil {
			return err
		}
	}

	return nil
}
