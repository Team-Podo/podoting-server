package models

import "github.com/Team-Podo/podoting-server/repository"

type ScheduleRepository interface {
	FindByPerformanceID(performanceID uint) ([]repository.Schedule, error)
	Find(uuid string) *repository.Schedule
	FindByUUID(uuid string) (*repository.Schedule, error)
	Save(schedule *repository.Schedule) error
	SaveMany(schedules []repository.Schedule) error
	Update(schedule *repository.Schedule) error
	Delete(uuid string) error
}
