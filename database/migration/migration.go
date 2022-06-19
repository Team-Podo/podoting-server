package migration

import (
	"github.com/kwanok/podonine/repository"
)

func Init() {
	db := repository.Gorm

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
