package models

type Performance interface {
	GetId() uint
	GetTitle() string
	GetStartDate() string
	GetEndDate() string
	GetProduct() Product
}

type PerformanceRepository interface {
	Get(query map[string]any) []Performance
	Find(id uint) Performance
	Save(performance Performance) Performance
	Update(performance Performance) Performance
	Delete(id uint)
}
