package models

import "github.com/Team-Podo/podoting-server/repository"

type ScheduleRepository interface {
	Get(query map[string]any) []repository.Schedule
	Find(uuid string) *repository.Schedule
	Save(schedule *repository.Schedule) error
	SaveMany(schedules []repository.Schedule) error
	Update(schedule *repository.Schedule) error
	Delete(uuid string) error
	GetTotal(query map[string]any) int64
}
