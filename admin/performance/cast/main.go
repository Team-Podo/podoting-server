package cast

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/repository"
)

func init() {
	repositories = Repository{
		cast:        &repository.CastRepository{DB: database.Gorm},
		performance: &repository.PerformanceRepository{DB: database.Gorm},
	}
}
