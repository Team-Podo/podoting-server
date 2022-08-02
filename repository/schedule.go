package repository

import (
	"database/sql"
	"errors"
	"github.com/Team-Podo/podoting-server/models"
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

func (s *Schedule) GetUUID() string {
	return s.UUID
}

func (s *Schedule) GetPerformance() models.Performance {
	return s.Performance
}

func (s *Schedule) GetMemo() string {
	return s.Memo
}

func (s *Schedule) GetDate() string {
	return s.Date
}

func (s *Schedule) GetTime() string {
	return s.Time.String
}

func (s *Schedule) GetCreatedAt() string {
	return s.CreatedAt.Format("2006-01-02 15:04:05")
}

func (s *Schedule) GetUpdatedAt() string {
	return s.UpdatedAt.Format("2006-01-02 15:04:05")
}

type ScheduleRepository struct {
	Db *gorm.DB
}

func (s *ScheduleRepository) Get(query map[string]any) []models.Schedule {
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

	var m = make([]models.Schedule, len(schedules))

	for i, schedule := range schedules {
		m[i] = &schedule
	}

	return m
}

func (s *ScheduleRepository) Find(uuid string) models.Schedule {
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

func (s *ScheduleRepository) Save(scheduleModel models.Schedule) models.Schedule {
	scheduleUUID, _ := uuid.NewUUID()

	schedule := Schedule{
		UUID:          scheduleUUID.String(),
		PerformanceId: scheduleModel.GetPerformance().GetId(),
		Memo:          scheduleModel.GetMemo(),
		Date:          scheduleModel.GetDate(),
		Time:          sql.NullString{String: scheduleModel.GetTime()},
	}

	result := s.Db.Create(&schedule)

	if result.Error != nil {
		return nil
	}

	return &schedule
}

func (s *ScheduleRepository) SaveMany(scheduleModels []models.Schedule) error {
	var schedules []Schedule

	for i := range scheduleModels {
		scheduleUUID, _ := uuid.NewUUID()
		scheduleModel := scheduleModels[i]
		schedules = append(schedules, Schedule{
			UUID: scheduleUUID.String(),
			Memo: scheduleModel.GetMemo(),
			Date: scheduleModel.GetDate(),
			Time: sql.NullString{String: scheduleModel.GetTime()},
		})
	}

	result := s.Db.Create(&schedules)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *ScheduleRepository) Update(scheduleModel models.Schedule) models.Schedule {
	err := s.Db.First(&Schedule{
		UUID: scheduleModel.GetUUID(),
	}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	s.Db.Model(&Schedule{UUID: scheduleModel.GetUUID()}).Updates(Schedule{
		UUID: scheduleModel.GetUUID(),
		Memo: scheduleModel.GetMemo(),
		Date: scheduleModel.GetDate(),
		Time: sql.NullString{String: scheduleModel.GetTime()},
	})

	return scheduleModel
}

func (s *ScheduleRepository) Delete(uuid string) {
	schedule := Schedule{}

	s.Db.Delete(&schedule)
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
