package cast

import (
	"github.com/Team-Podo/podoting-server/repository"
)

var repositories Repository

type Repository struct {
	cast        *repository.CastRepository
	performance *repository.PerformanceRepository
}
