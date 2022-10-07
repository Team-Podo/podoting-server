package seat_grade_get

import "github.com/Team-Podo/podoting-server/repository"

type SeatGrade struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Color     string `json:"color"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ParseResponseForm(seatGrades []repository.SeatGrade) []SeatGrade {
	var response []SeatGrade
	for _, seatGrade := range seatGrades {
		response = append(response, SeatGrade{
			ID:        seatGrade.ID,
			Name:      seatGrade.Name,
			Price:     seatGrade.Price,
			Color:     seatGrade.Color,
			CreatedAt: seatGrade.CreatedAt.String(),
			UpdatedAt: seatGrade.UpdatedAt.String(),
		})
	}

	return response
}
