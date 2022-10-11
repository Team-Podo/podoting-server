package models

import "github.com/Team-Podo/podoting-server/repository"

type PerformanceContentRepository interface {
	Save(content *repository.PerformanceContent) error
	FindOneByID(id uint) *repository.PerformanceContent
}
