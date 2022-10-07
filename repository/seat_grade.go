package repository

import (
	"gorm.io/gorm"
	"time"
)

type SeatGrade struct {
	ID            uint            `json:"id" gorm:"primaryKey"`
	Performance   *Performance    `json:"performance" gorm:"foreignKey:PerformanceID"`
	PerformanceID uint            `json:"-"`
	Name          string          `json:"name"`
	Price         int             `json:"price"`
	Color         string          `json:"color"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"index"`
}

type SeatGradeRepository struct {
	DB *gorm.DB
}

func (r *SeatGradeRepository) GetAll() ([]SeatGrade, error) {
	var seatGrades []SeatGrade
	err := r.DB.Find(seatGrades).Error
	if err != nil {
		return nil, err
	}

	return seatGrades, nil
}

func (r *SeatGradeRepository) GetByID(id uint) (*SeatGrade, error) {
	seatGrade := &SeatGrade{}
	err := r.DB.First(seatGrade, id).Error
	if err != nil {
		return nil, err
	}

	return seatGrade, nil
}

func (r *SeatGradeRepository) GetByPerformanceID(performanceID uint) ([]SeatGrade, error) {
	var seatGrades []SeatGrade
	err := r.DB.
		Where("performance_id = ?", performanceID).
		Find(&seatGrades).Error
	if err != nil {
		return nil, err
	}

	return seatGrades, nil
}

func (r *SeatGradeRepository) Create(seatGrade *SeatGrade) error {
	err := r.DB.Create(seatGrade).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *SeatGradeRepository) Update(seatGrade *SeatGrade) error {
	err := r.DB.First(&SeatGrade{}, seatGrade.ID).Error
	if err != nil {
		return err
	}

	err = r.DB.Model(&SeatGrade{ID: seatGrade.ID}).
		Updates(seatGrade).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *SeatGradeRepository) Delete(seatGrade *SeatGrade) error {
	err := r.DB.Delete(seatGrade).Error
	if err != nil {
		return err
	}

	return nil
}
