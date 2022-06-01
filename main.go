package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/repository"
	"github.com/kwanok/podonine/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sqlite/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect gorm")
	}

	err = db.AutoMigrate(&repository.Place{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&repository.Product{})
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

	repository.Init()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.Routes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
