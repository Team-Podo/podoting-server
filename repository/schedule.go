package repository

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

type Schedule struct {
	UUID          string `json:"uuid" gorm:"primarykey"`
	Performance   *Performance
	PerformanceID uint `json:"-"`
	Memo          string
	Date          string
	Open          bool   `json:"open" gorm:"default:false"`
	Casts         []Cast `gorm:"many2many:schedule_cast;"`
	Time          sql.NullString
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"index"`
}

type ScheduleRepository struct {
	DB *gorm.DB
}

func (s *ScheduleRepository) FindByPerformanceID(performanceID uint) ([]Schedule, error) {
	var schedules []Schedule

	err := s.DB.
		Preload("Casts").
		Preload("Casts.Person").
		Preload("Casts.Character").
		Preload("Casts.ProfileImage").
		Where("performance_id = ?", performanceID).
		Find(&schedules).Error

	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (s *ScheduleRepository) Find(uuid string) *Schedule {
	schedule := Schedule{
		UUID: uuid,
	}

	err := s.DB.Preload("Performance").First(&schedule).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return &schedule
}

func (s *ScheduleRepository) FindByUUID(uuids string) (*Schedule, error) {
	var schedule Schedule

	err := s.DB.Where("uuid = ?", uuids).First(&schedule).Error

	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (s *ScheduleRepository) Save(schedule *Schedule) error {
	scheduleUUID, _ := uuid.NewUUID()

	schedule.UUID = scheduleUUID.String()

	result := s.DB.Create(&schedule)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ScheduleRepository) SaveMany(schedules []Schedule) error {
	err := s.DB.Create(&schedules).Error

	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRepository) Update(schedule *Schedule) error {
	err := s.DB.Unscoped().Model(&schedule).Association("Casts").Replace(schedule.Casts)
	if err != nil {
		return errors.New("failed to update casts")
	}

	err = s.DB.Model(&Schedule{UUID: schedule.UUID}).
		Updates(schedule).Error

	if err != nil {
		return errors.New("failed to update schedule")
	}

	return nil
}

func (s *ScheduleRepository) Delete(uuid string) error {
	schedule := Schedule{UUID: uuid}

	err := s.DB.Delete(&schedule).Error
	if err != nil {
		return err
	}

	return nil
}
