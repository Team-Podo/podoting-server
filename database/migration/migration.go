package migration

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/repository"
)

func Init() {
	db := database.Gorm

	err := db.AutoMigrate(
		&repository.AreaBoilerplate{},
		&repository.PerformanceContent{},
		&repository.Point{},
		&repository.Seat{},
		&repository.SeatBooking{},
		&repository.SeatGrade{},
		&repository.Location{},
		&repository.Place{},
		&repository.Schedule{},
		&repository.Performance{},
		&repository.Product{},
		&repository.File{},
		&repository.ProductContent{},
		&repository.Person{},
		&repository.Character{},
		&repository.Cast{},
		&repository.PerformanceCast{},
	)
	if err != nil {
		panic(err)
	}
}
