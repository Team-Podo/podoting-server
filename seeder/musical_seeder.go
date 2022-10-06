package seeder

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/google/uuid"
	"strconv"
)

func SeedMusical11() {
	db := database.Gorm
	performance := repository.Performance{ID: 11}
	db.Debug().
		Preload("Areas.Seats").
		Find(&performance)
	fmt.Println(performance.Areas[0].Seats)

	seatSeeder()

	db.Debug().
		Preload("Areas.Seats").
		Find(&performance)
	fmt.Println(performance.Areas[0].Seats)
}

func getMusical11() repository.Performance {
	return repository.Performance{
		ID: 11,
	}
}

func seatSeeder() {
	seatRepository := repository.SeatRepository{DB: database.Gorm}
	_ = seatRepository.CreateSeats(getSeats())
}

func getSeats() []repository.Seat {
	var seats []repository.Seat
	for i := 0; i < 100; i++ {
		seat := repository.Seat{
			UUID:   uuid.New().String(),
			Name:   "1열 " + strconv.Itoa(i+1) + "번",
			AreaID: 4,
			Grade:  &repository.SeatGrade{ID: 1},
			Point: &repository.Point{
				X: float64(i),
				Y: float64(i / 10),
			},
		}
		seats = append(seats, seat)
	}
	return seats
}
