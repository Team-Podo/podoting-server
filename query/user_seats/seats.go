package user_seats

type Seat struct {
	UUID        string  `json:"uuid"`
	SeatName    string  `json:"seatName"`
	GradeName   string  `json:"gradeName"`
	Price       int     `json:"price"`
	Color       string  `json:"color"`
	PointX      float64 `json:"pointX"`
	PointY      float64 `json:"pointY"`
	BookedCount int     `json:"bookedCount"`
}
