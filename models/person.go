package models

import "github.com/Team-Podo/podoting-server/repository"

type PersonRepository interface {
	FindAll() ([]repository.Person, error)
	FindByID(id uint) (*repository.Person, error)
	Create(person *repository.Person) error
	Update(person *repository.Person) error
	Delete(id uint) error
}
