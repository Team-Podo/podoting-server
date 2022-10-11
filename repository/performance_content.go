package repository

import (
	"gorm.io/gorm"
	"time"
)

type PerformanceContent struct {
	ID            uint         `json:"id" gorm:"primaryKey"`
	ManagingTitle string       `json:"managingTitle"`
	Content       string       `json:"content" gorm:"type:text"`
	Visible       bool         `json:"visible"`
	Performance   *Performance `json:"performance" gorm:"foreignKey:PerformanceID"`
	PerformanceID uint         `json:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"index"`
}

type PerformanceContentRepository struct {
	DB *gorm.DB
}

func (p *PerformanceContentRepository) Save(content *PerformanceContent) error {
	err := p.DB.Save(&content).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *PerformanceContentRepository) FindOneByID(id uint) *PerformanceContent {
	var content PerformanceContent
	err := p.DB.First(&content, id).Error
	if err != nil {
		return nil
	}

	return &content
}
