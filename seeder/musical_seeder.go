package seeder

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/google/uuid"
)

func SeedMusical11() {
	seatSeeder()
}

func seatSeeder() {
	seatRepository := repository.SeatRepository{DB: database.Gorm}
	_ = seatRepository.CreateSeats(getSeats())
}

func getSeats() []repository.Seat {
	var seats []repository.Seat
	for i := 0; i < 112; i++ {
		seat := repository.Seat{
			UUID: uuid.New().String(),
			AreaBoilerplate: &repository.AreaBoilerplate{
				AreaID: 26,
				Point: &repository.Point{
					X: 0,
					Y: 0,
				},
			},
			SeatGradeID:   1,
			PerformanceID: 1,
		}
		seats = append(seats, seat)
	}
	return seats
}
