package models

import "github.com/Team-Podo/podoting-server/repository"

type OrderRepository interface {
	Save(order *repository.Order) error
	GetByUserUID(userUID string) []repository.Order
	GetByUserUIDWithQuery(userUID string, query map[string]any) ([]repository.Order, int64)
	FindByID(id uint) *repository.Order
	CancelOrder(order *repository.Order) error
}
