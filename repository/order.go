package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/utils"
	"log"
	"strconv"
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
		Joins("Schedule").
		Where("buyer_uid = ?", userUID).
		Find(&orders)

	return orders
}

func (r *OrderRepository) GetByUserUIDWithQuery(userUID string, query map[string]any) ([]Order, int64) {
	var orders []Order

	limit, _ := strconv.Atoi(utils.ToString(query["limit"]))
	offset, _ := strconv.Atoi(utils.ToString(query["offset"]))
	reversed, _ := query["reversed"].(bool)

	r.DB.
		Limit(limit).
		Offset(offset).
		Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: reversed}).
		Preload("Details.SeatBooking.Seat.Grade").
		Preload("Details.SeatBooking.Seat.AreaBoilerplate").
		Preload("Performance.Thumbnail").
		Preload("Performance.Place").
		Joins("Schedule").
		Where("buyer_uid = ?", userUID).
		Find(&orders)

	var count int64
	r.DB.
		Model(&Order{}).
		Where("buyer_uid = ?", userUID).
		Count(&count)

	return orders, count
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

	var seatBookingIds []uint

	for _, detail := range order.Details {
		seatBookingIds = append(seatBookingIds, detail.SeatBookingID)
	}

	now := time.Now()

	err := r.DB.Model(Order{}).
		Where("id = ?", order.ID).
		Updates(Order{
			Canceled:   true,
			CanceledAt: &now,
		}).Error

	if err != nil {
		return err
	}

	err = r.DB.Model(OrderDetail{}).
		Where("order_id = ?", order.ID).
		Updates(OrderDetail{
			Canceled:   true,
			CanceledAt: &now,
		}).Error

	if err != nil {
		return err
	}

	err = r.DB.Model(SeatBooking{}).
		Where("id IN ?", seatBookingIds).
		Updates(SeatBooking{
			Booked:     false,
			Canceled:   true,
			CanceledAt: &now,
		}).Error

	if err != nil {
		return err
	}

	return nil
}
