package models

import "github.com/Team-Podo/podoting-server/repository"

type CastRepository interface {
	Get() ([]repository.Cast, error)
	FindByPerformanceID(performanceID uint) ([]repository.Cast, error)
	FindOneByID(id uint) (*repository.Cast, error)
	CreateMany(casts []repository.Cast) error
	SavePerformanceCasts(performanceCasts []repository.PerformanceCast) error
	Update(cast *repository.Cast) error
	Delete(id uint) error
}
