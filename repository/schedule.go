package repository

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"log"
	"strconv"
	"time"
)

type Schedule struct {
	UUID          string `json:"uuid" gorm:"primarykey"`
	Performance   *Performance
	PerformanceId uint `json:"-"`
	Memo          string
	Date          string
	Time          sql.NullString
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"index"`
}

type ScheduleRepository struct {
	Db *gorm.DB
}

func (s *ScheduleRepository) Get(query map[string]any) []Schedule {
	var schedules []Schedule
	db := s.Db

	if query["reversed"] == true {
		db = db.Order("id desc")
	}

	if query["limit"] != nil {
		limit, _ := strconv.Atoi(utils.ToString(query["limit"]))
		db = db.Limit(limit)
	}

	if query["offset"] != nil {
		offset, _ := strconv.Atoi(utils.ToString(query["offset"]))
		db = db.Offset(offset)
	}

	err := db.Preload("Performance").Find(&schedules).Error

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	if len(schedules) == 0 {
		return nil
	}

	return schedules
}

func (s *ScheduleRepository) Find(uuid string) *Schedule {
	schedule := Schedule{
		UUID: uuid,
	}

	err := s.Db.Preload("performance").First(&schedule).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return &schedule
}

func (s *ScheduleRepository) Save(schedule *Schedule) error {
	scheduleUUID, _ := uuid.NewUUID()

	schedule.UUID = scheduleUUID.String()

	result := s.Db.Create(&schedule)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ScheduleRepository) SaveMany(schedules []Schedule) error {
	err := s.Db.Create(&schedules).Error

	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRepository) Update(schedule *Schedule) error {
	err := s.Db.First(&Schedule{
		UUID: schedule.UUID,
	}).Error

	if err != nil {
		return err
	}

	err = s.Db.Model(&Schedule{UUID: schedule.UUID}).Updates(schedule).Error

	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRepository) Delete(uuid string) error {
	schedule := Schedule{UUID: uuid}

	err := s.Db.Delete(&schedule).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduleRepository) GetTotal(query map[string]any) int64 {
	db := s.Db

	if query["reversed"] == true {
		db = db.Order("id desc")
	}

	if query["limit"] != nil {
		limit, _ := strconv.Atoi(utils.ToString(query["limit"]))
		db = db.Limit(limit)
	}

	if query["offset"] != nil {
		offset, _ := strconv.Atoi(utils.ToString(query["offset"]))
		db = db.Offset(offset)
	}

	var count int64
	db.Model(&Schedule{}).Count(&count)

	return count
}
