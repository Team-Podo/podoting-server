package performance

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
)

var repositories Repository

type request struct {
	PlaceID     uint   `json:"placeID"`
	MainAreaID  uint   `json:"mainAreaID"`
	Title       string `json:"title"`
	RunningTime string `json:"runningTime"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Rating      string `json:"rating"`
}

type Repository struct {
	performance     models.PerformanceRepository
	file            models.FileRepository
	areaBoilerplate models.AreaBoilerplateRepository
}

func init() {
	repositories = Repository{
		performance: &repository.PerformanceRepository{DB: database.Gorm},
		file:        &repository.FileRepository{DB: database.Gorm},
		areaBoilerplate: &repository.AreaBoilerplateRepository{
			DB: database.Gorm,
		},
	}
}
