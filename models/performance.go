package models

type Performance interface {
	GetId() uint
	GetTitle() string
	GetStartDate() string
	GetEndDate() string
}

type PerformanceRepository interface {
	Get() []Performance
	Find(id uint) Performance
	Save(performance Performance) Performance
	Update(performance Performance) Performance
	Delete(id uint)
}
