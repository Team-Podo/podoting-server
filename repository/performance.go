package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"strconv"
	"time"
)

type Performance struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	ProductID uint            `json:"-"`
	Product   *Product        `json:"product" gorm:"foreignkey:ProductID"`
	Schedules []Schedule      `gorm:"foreignkey:PerformanceId"`
	Title     string          `json:"title"`
	StartDate string          `json:"startDate"`
	EndDate   string          `json:"endDate"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}

type PerformanceRepository struct {
	Db *gorm.DB
}

func (p *PerformanceRepository) GetWithQueryMap(query map[string]any) []Performance {
	var performances []Performance
	db := p.Db

	p.applyAllQuery(query)

	err := db.Find(&performances).Error

	if err != nil {
		return nil
	}

	if len(performances) == 0 {
		return nil
	}

	return performances
}

func (p *PerformanceRepository) FindByID(id uint) *Performance {
	performance := Performance{}
	err := p.Db.Preload("Product").Preload("Schedules").First(&performance, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		return nil
	}

	return &performance
}

func (p *PerformanceRepository) Save(performance *Performance) error {
	err := p.Db.Create(performance).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *PerformanceRepository) Update(performance *Performance) error {
	err := p.Db.First(&Performance{}, performance.ID).Error

	if err != nil {
		return err
	}

	p.Db.Model(&Performance{ID: performance.ID}).Updates(performance)

	return nil
}

func (p *PerformanceRepository) Delete(id uint) error {
	performance := Performance{}
	performance.ID = id

	err := p.Db.Delete(&performance).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *PerformanceRepository) GetTotalWithQueryMap(query map[string]any) int64 {
	p.applyAllQuery(query)

	var count int64
	p.Db.Model(&Performance{}).Count(&count)

	return count
}

func (p *PerformanceRepository) applyAllQuery(query map[string]any) {
	p.applyReversedQuery(query)
	p.applyLimitQuery(query)
	p.applyOffsetQuery(query)
}

func (p *PerformanceRepository) applyReversedQuery(query map[string]any) {
	if query["reversed"] == true {
		p.Db = p.Db.Order("id desc")
	}
}

func (p *PerformanceRepository) applyLimitQuery(query map[string]any) {
	if query["limit"] != nil {
		limit, _ := strconv.Atoi(utils.ToString(query["limit"]))
		p.Db = p.Db.Limit(limit)
	}
}

func (p *PerformanceRepository) applyOffsetQuery(query map[string]any) {
	if query["offset"] != nil {
		offset, _ := strconv.Atoi(utils.ToString(query["offset"]))
		p.Db = p.Db.Offset(offset)
	}
}
