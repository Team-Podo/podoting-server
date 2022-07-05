package database

import (
	"fmt"
	"github.com/kwanok/podonine/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var Gorm = initDB()

func initDB() *gorm.DB {
	utils.SetEnv()

	connection := os.Getenv("DB_CONNECTION")

	fmt.Println("커넥션:", connection)

	if connection == "mysql" {
		dsn := "root:password@tcp(127.0.0.1:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect gorm")
		}

		return db
	} else if connection == "sqlite" {
		db, err := gorm.Open(sqlite.Open(utils.RootPath()+"sqlite/test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect gorm")
		}

		return db
	} else {
		return nil
	}

}
