package models

import "github.com/Team-Podo/podoting-server/repository"

type OrderDetailRepository interface {
	Save(order *repository.OrderDetail) error
	FindByID(id uint) *repository.OrderDetail
	CancelOrderDetail(order *repository.OrderDetail) error
}
