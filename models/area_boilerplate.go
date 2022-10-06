package models

import "github.com/Team-Podo/podoting-server/repository"

type AreaBoilerplateRepository interface {
	GetAreaBoilerplateByAreaID(areaID uint) ([]repository.AreaBoilerplate, error)
	SaveSeats(boilerplates []repository.AreaBoilerplate, performanceID uint) error
}
