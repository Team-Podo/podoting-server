package models

import "github.com/Team-Podo/podoting-server/repository"

type PerformanceRepository interface {
	GetWithQueryMap(query map[string]any) []repository.Performance
	FindByID(id uint) *repository.Performance
	Save(performance *repository.Performance) error
	Update(performance *repository.Performance) error
	Delete(id uint) error
	GetTotalWithQueryMap(query map[string]any) int64
}
