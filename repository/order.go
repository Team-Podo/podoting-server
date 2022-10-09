package repository

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID            uint            `json:"id" gorm:"primary_key"`
	OrderKey      string          `json:"order_key" gorm:"unique_index;not null;size:8"`
	Performance   *Performance    `json:"performance" gorm:"foreignKey:PerformanceID"`
	PerformanceID uint            `json:"-"`
	BuyerUID      string          `json:"buyerUID"`
	Details       []OrderDetail   `json:"details" gorm:"foreignKey:OrderID"`
	Paid          bool            `json:"paid" gorm:"default:false"`
	Canceled      bool            `json:"canceled" gorm:"default:false"`
	CanceledAt    *time.Time      `json:"canceledAt"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"index"`
}

type OrderRepository struct {
	DB *gorm.DB
}

func (r *OrderRepository) Save(order *Order) error {
	err := r.DB.Save(order).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetByUserUID(userUID string) []Order {
	var orders []Order
	r.DB.
		Preload("Details.SeatBooking.Seat.Grade").
		Preload("Details.SeatBooking.Seat.AreaBoilerplate").
		Preload("Performance.Thumbnail").
		Where("buyer_uid = ?", userUID).
		Find(&orders)

	return orders
}
