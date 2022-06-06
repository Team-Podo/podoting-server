package repository

import (
	"fmt"
	"github.com/kwanok/podonine/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var Gorm *gorm.DB

type Model struct {
	ID        uint            `gorm:"primarykey"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}

func init() {
	db, err := gorm.Open(sqlite.Open(utils.RootPath()+"sqlite/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect gorm")
	}

	Gorm = db

	fmt.Println("[GORM] Connected!")
}
