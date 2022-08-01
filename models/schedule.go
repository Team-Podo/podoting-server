package models

type Schedule interface {
	GetUUID() string
	GetPerformance() Performance
	GetMemo() string
	GetDate() string
	GetTime() string
	GetCreatedAt() string
	GetUpdatedAt() string
}

type ScheduleRepository interface {
	Get(query map[string]any) []Schedule
	Find(uuid string) Schedule
	Save(schedule Schedule) Schedule
	SaveMany(schedules []Schedule) error
	Update(schedule Schedule) Schedule
	Delete(uuid string)
	GetTotal(query map[string]any) int64
}
