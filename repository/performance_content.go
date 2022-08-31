package repository

import (
	"gorm.io/gorm"
	"time"
)

type PerformanceContent struct {
	UUID          string       `json:"uuid" gorm:"primarykey"`
	Performance   *Performance `json:"performance"`
	PerformanceID uint         `json:"-"`
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	Content       string       `json:"content"`
	Priority      uint         `json:"priority"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"index"`
}

type PerformanceContentRepository struct {
	DB *gorm.DB
}

func (p *PerformanceContentRepository) Save(content *PerformanceContent) error {
	err := p.DB.Create(content).Error
	if err != nil {
		return err
	}

	return nil
}
