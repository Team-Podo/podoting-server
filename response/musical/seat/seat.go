package seat

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
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
