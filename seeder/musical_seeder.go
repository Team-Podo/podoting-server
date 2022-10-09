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
	for i := 5; i < 100; i++ {
		ui := uint(i)
		seat := repository.Seat{
			UUID:              uuid.New().String(),
			AreaBoilerplateID: &ui,
			SeatGradeID:       1,
			PerformanceID:     1,
		}
		seats = append(seats, seat)
	}
	return seats
}
