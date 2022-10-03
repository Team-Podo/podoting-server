package models

import "github.com/Team-Podo/podoting-server/repository"

type CastRepository interface {
	Get() ([]repository.Cast, error)
	GetByPerformanceID(performanceID uint) ([]repository.Cast, error)
	FindByID(id uint) (*repository.Cast, error)
	Create(cast *repository.Cast) error
	Update(cast *repository.Cast) error
	Delete(id uint) error
}
