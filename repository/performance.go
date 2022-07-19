package repository

import (
	"errors"
	"github.com/kwanok/podonine/models"
	"gorm.io/gorm"
	"log"
)

type Performance struct {
	Model
	Title     string
	StartDate string
	EndDate   string
}

func (p *Performance) GetId() uint {
	return p.ID
}

func (p *Performance) GetTitle() string {
	return p.Title
}

func (p *Performance) GetStartDate() string {
	return p.StartDate
}

func (p *Performance) GetEndDate() string {
	return p.EndDate
}

type PerformanceRepository struct {
	Db *gorm.DB
}

func (p *PerformanceRepository) Get() []models.Performance {
	var performances []*Performance
	err := p.Db.Find(&performances).Error

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	if len(performances) == 0 {
		return nil
	}

	var m = make([]models.Performance, len(performances))
	for i, performance := range performances {
		m[i] = performance
	}

	return m
}

func (p *PerformanceRepository) Find(id uint) models.Performance {
	performance := Performance{}
	err := p.Db.First(&performance, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return &performance
}

func (p *PerformanceRepository) Save(m models.Performance) models.Performance {
	performance := Performance{
		Title:     m.GetTitle(),
		StartDate: m.GetStartDate(),
		EndDate:   m.GetEndDate(),
	}

	result := p.Db.Create(&performance)

	if result.Error != nil {
		return nil
	}

	return &performance
}

func (p *PerformanceRepository) Update(m models.Performance) models.Performance {
	performance := Performance{}
	performance.ID = m.GetId()

	tx := p.Db.First(&performance)

	performance.Title = m.GetTitle()
	performance.StartDate = m.GetStartDate()
	performance.EndDate = m.GetEndDate()

	tx.Save(&performance)

	return &performance
}

func (p *PerformanceRepository) Delete(id uint) {
	performance := Performance{}
	performance.ID = id

	p.Db.Delete(&performance)
}
