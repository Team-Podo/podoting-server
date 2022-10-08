package seat

import (
	"database/sql"
	"github.com/Team-Podo/podoting-server/repository"
)

type Seat struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Point  Point  `json:"point"`
	Grade  Grade  `json:"grade"`
	Booked bool   `json:"booked"`
	Color  string `json:"color"`
	Price  int    `json:"price"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Grade struct {
	Name string `json:"name"`
}

type Schedule struct {
	UUID string `json:"uuid"`
	Date string `json:"date"`
	Time string `json:"time"`
}

func ParseSeats(seats []repository.Seat) []Seat {
	response := make([]Seat, len(seats))

	for i := range seats {
		Booked := false
		s := seats[i]

		if len(seats[i].Bookings) > 0 {
			Booked = true
		}

		response[i] = Seat{
			UUID: s.UUID,
			Name: s.AreaBoilerplate.Name,
			Point: Point{
				X: s.AreaBoilerplate.Point.X,
				Y: s.AreaBoilerplate.Point.Y,
			},
			Grade:  Grade{Name: s.Grade.Name},
			Booked: Booked,
			Color:  s.Grade.Color,
			Price:  s.Grade.Price,
		}
	}

	return response
}

func ParseSchedules(schedules []repository.Schedule) []Schedule {
	response := make([]Schedule, len(schedules))

	for i := range schedules {
		response[i] = Schedule{
			UUID: schedules[i].UUID,
			Date: schedules[i].Date,
			Time: func(time sql.NullString) string {
				if time.Valid {
					return time.String
				}
				return ""
			}(schedules[i].Time),
		}
	}

	return response
}
