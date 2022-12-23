package mypage

import (
	"github.com/Team-Podo/podoting-server/repository"
)

type Order struct {
	ID            uint          `json:"id"`
	Key           string        `json:"key"`
	Performance   Performance   `json:"performance"`
	Schedule      Schedule      `json:"schedule"`
	PaymentMethod string        `json:"paymentMethod"`
	Canceled      bool          `json:"canceled"`
	Details       []OrderDetail `json:"details"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
}

type Performance struct {
	ID        uint    `json:"id"`
	Title     string  `json:"title"`
	Place     string  `json:"place"`
	Thumbnail *string `json:"thumbnail"`
}

type OrderDetail struct {
	ID            uint    `json:"id"`
	Key           string  `json:"key"`
	Canceled      bool    `json:"canceled"`
	CanceledAt    *string `json:"canceledAt"`
	OriginalPrice int     `json:"originalPrice"`
	DiscountPrice int     `json:"discountPrice"`
	PayPrice      int     `json:"payPrice"`
	Seat          Seat    `json:"seat"`
}

type Schedule struct {
	UUID string  `json:"uuid"`
	Date string  `json:"date"`
	Time *string `json:"time"`
}

type Seat struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Grade string `json:"grade"`
	Color string `json:"color"`
	Price int    `json:"price"`
}

func ParseOrder(orders []repository.Order) []Order {
	var response = make([]Order, len(orders))
	for i := range response {
		response[i] = Order{
			ID:          orders[i].ID,
			Canceled:    orders[i].Canceled,
			Key:         orders[i].OrderKey,
			Performance: ParsePerformance(orders[i].Performance),
			Schedule:    ParseSchedule(orders[i].Schedule),
			Details:     ParseOrderDetail(orders[i].Details),
			CreatedAt:   orders[i].CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   orders[i].UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return response
}

func ParsePerformance(performance *repository.Performance) Performance {
	return Performance{
		ID:    performance.ID,
		Title: performance.Title,
		Place: performance.Place.Name,
		Thumbnail: func() *string {
			if performance.Thumbnail != nil {
				fullPath := performance.Thumbnail.FullPath()
				return &fullPath
			}
			return nil
		}(),
	}
}

func ParseSchedule(schedule *repository.Schedule) Schedule {
	return Schedule{
		UUID: schedule.UUID,
		Date: schedule.Date,
		Time: func() *string {
			if schedule.Time.Valid {
				return &schedule.Time.String
			}
			return nil
		}(),
	}
}

func ParseOrderDetail(details []repository.OrderDetail) []OrderDetail {
	var response = make([]OrderDetail, len(details))
	for i := range response {
		response[i] = OrderDetail{
			ID:       details[i].ID,
			Key:      details[i].OrderDetailKey,
			Canceled: details[i].Canceled,
			CanceledAt: func() *string {
				if details[i].CanceledAt != nil {
					canceledAt := details[i].CanceledAt.Format("2006-01-02 15:04:05")
					return &canceledAt
				}
				return nil
			}(),
			OriginalPrice: int(details[i].OriginalPrice),
			DiscountPrice: int(details[i].Discount),
			PayPrice:      int(details[i].OriginalPrice - details[i].Discount),
			Seat: Seat{
				UUID:  details[i].SeatBooking.Seat.UUID,
				Name:  details[i].SeatBooking.Seat.AreaBoilerplate.Name,
				Grade: details[i].SeatBooking.Seat.Grade.Name,
				Color: details[i].SeatBooking.Seat.Grade.Color,
				Price: details[i].SeatBooking.Seat.Grade.Price,
			},
		}
	}

	return response
}
