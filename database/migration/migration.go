package migration

import (
	"github.com/kwanok/podonine/database"
	"github.com/kwanok/podonine/repository"
)

func Init() {
	db := database.Gorm

	err := db.AutoMigrate(&repository.Place{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&repository.Product{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&repository.ProductContent{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&repository.Seat{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&repository.Area{})
	if err != nil {
		return
	}
}
