package models

import "github.com/Team-Podo/podoting-server/repository"

type OrderRepository interface {
	Save(order *repository.Order) error
	GetByUserUID(userUID string) []repository.Order
}
