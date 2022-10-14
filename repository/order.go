package repository

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Order struct {
	ID            uint            `json:"id" gorm:"primary_key"`
	OrderKey      string          `json:"order_key" gorm:"unique_index;not null;size:8"`
	Performance   *Performance    `json:"performance" gorm:"foreignKey:PerformanceID"`
	PerformanceID uint            `json:"-"`
	Schedule      *Schedule       `json:"schedule" gorm:"foreignKey:ScheduleUUID"`
	ScheduleUUID  string          `json:"-" gorm:"size:36"`
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
	err := r.DB.Debug().Save(order).Error
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
		Joins("Schedule").
		Where("buyer_uid = ?", userUID).
		Find(&orders)

	return orders
}

func (r *OrderRepository) FindByID(ID uint) *Order {
	var order *Order
	err := r.DB.
		Preload("Details.SeatBooking").
		Joins("Schedule").
		First(&order, ID).
		Error

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return order
}

func (r *OrderRepository) CancelOrder(order *Order) error {
	err := r.DB.Model(order).
		Debug().
		Update("canceled", true).
		Update("canceled_at", time.Now()).
		Error

	details := order.Details

	err = r.DB.Model(&details).
		Debug().
		Update("canceled", true).
		Update("canceled_at", time.Now()).
		Error

	seatBookings := make([]SeatBooking, len(details))
	for i, detail := range details {
		seatBookings[i] = *detail.SeatBooking
	}

	err = r.DB.Model(&seatBookings).
		Debug().
		Update("booked", false).
		Update("canceled", true).
		Update("canceled_at", time.Now()).
		Error

	if err != nil {
		return err
	}

	return nil

}
