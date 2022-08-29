package models

import "github.com/Team-Podo/podoting-server/repository"

type CastRepository interface {
	FindByID(id uint) (*repository.Cast, error)
	Update(cast *repository.Cast) error
}
