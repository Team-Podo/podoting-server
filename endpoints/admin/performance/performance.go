package performance

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"strconv"
)

var repositories Repository

type request struct {
	ProductID   uint   `json:"productId"`
	PlaceID     uint   `json:"placeId"`
	Title       string `json:"title"`
	RunningTime string `json:"runningTime"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Rating      string `json:"rating"`
}

type Repository struct {
	performance models.PerformanceRepository
	file        models.FileRepository
}

func init() {
	repositories = Repository{
		performance: &repository.PerformanceRepository{DB: database.Gorm},
		file:        &repository.FileRepository{DB: database.Gorm},
	}
}

func parseUint(s string) (uint, error) {
	i, err := strconv.Atoi(s)
	return uint(i), err
}
