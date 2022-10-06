package seat_get

import "github.com/Team-Podo/podoting-server/repository"

type Seat struct {
	UUID  string `json:"uuid"`
	Grade Grade  `json:"grade"`
	Point Point  `json:"point"`
}

type Grade struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Color string `json:"color"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func ParseResponseForm(seats []repository.Seat) []Seat {
	var response []Seat
	for _, seat := range seats {
		response = append(response, Seat{
			UUID:  seat.UUID,
			Grade: Grade{ID: seat.Grade.ID, Name: seat.Grade.Name, Price: seat.Grade.Price, Color: seat.Grade.Color},
			Point: Point{X: seat.Point.X, Y: seat.Point.Y},
		})
	}

	return response
}
