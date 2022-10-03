package models

import "github.com/Team-Podo/podoting-server/repository"

type CharacterRepository interface {
	FindByPerformanceID(performanceID uint) ([]repository.Character, error)
	Create(character *repository.Character) error
	Update(character *repository.Character) error
	Delete(id uint) error
}
