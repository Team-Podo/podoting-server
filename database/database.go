package database

import (
	"github.com/kwanok/podonine/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Gorm = initDB()

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(utils.RootPath()+"sqlite/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect gorm")
	}

	return db
}
