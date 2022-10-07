package models

import "github.com/Team-Podo/podoting-server/repository"

type SeatGradeRepository interface {
	GetAll() ([]repository.SeatGrade, error)
	GetByID(id uint) (*repository.SeatGrade, error)
	GetByPerformanceID(performanceID uint) ([]repository.SeatGrade, error)
	Create(seatGrade *repository.SeatGrade) error
	Update(seatGrade *repository.SeatGrade) error
	Delete(seatGrade *repository.SeatGrade) error
}
