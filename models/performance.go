package models

type Performance interface {
	GetId() uint
	GetTitle() string
	GetStartDate() string
	GetEndDate() string
	GetProduct() Product
	GetSchedules() []Schedule
	GetCreatedAt() string
	GetUpdatedAt() string
}

type PerformanceRepository interface {
	Get(query map[string]any) []Performance
	Find(id uint) Performance
	Save(performance Performance) Performance
	Update(performance Performance) Performance
	Delete(id uint)
	GetTotal(query map[string]any) int64
}
