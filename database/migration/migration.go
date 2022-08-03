package migration

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/repository"
)

func Init() {
	db := database.Gorm

	err := db.AutoMigrate(&repository.Schedule{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&repository.Performance{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&repository.Product{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&repository.File{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&repository.ProductContent{})
	if err != nil {
		return
	}
}
