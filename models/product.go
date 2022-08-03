package models

import "github.com/Team-Podo/podoting-server/repository"

type ProductRepository interface {
	GetWithQueryMap(query map[string]any) ([]repository.Product, error)
	FindByID(id uint) (*repository.Product, error)
	Save(product *repository.Product) error
	Update(product *repository.Product) error
	Delete(id uint) error
	GetTotalWithQueryMap(query map[string]any) (int64, error)
}
