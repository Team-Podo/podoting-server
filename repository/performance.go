package repository

import (
	"errors"
	"github.com/kwanok/podonine/models"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"log"
	"strconv"
)

type Performance struct {
	Model
	Product   *Product `gorm:"foreignkey:ProductID"`
	ProductID uint     `json:"-"`
	Title     string
	StartDate string
	EndDate   string
}

func (p *Performance) GetId() uint {
	return p.ID
}

func (p *Performance) GetProduct() models.Product {
	return p.Product
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

func (p *PerformanceRepository) Get(query map[string]any) []models.Performance {
	var performances []*Performance
	db := p.Db

	if query["reversed"] == true {
		db = db.Order("id desc")
	}

	if query["limit"] != nil {
		limit, _ := strconv.Atoi(utils.ToString(query["limit"]))
		db = db.Limit(limit)
	}

	if query["offset"] != nil {
		offset, _ := strconv.Atoi(utils.ToString(query["offset"]))
		db = db.Offset(offset)
	}

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
	err := p.Db.Preload("Product").First(&performance, id).Error

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
		ProductID: m.GetProduct().GetId(),
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
	err := p.Db.First(&Performance{}, m.GetId()).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	p.Db.Model(&Performance{Model: Model{ID: m.GetId()}}).Updates(Performance{
		ProductID: m.GetProduct().GetId(),
		Title:     m.GetTitle(),
		StartDate: m.GetStartDate(),
		EndDate:   m.GetEndDate(),
	})

	return m
}

func (p *PerformanceRepository) Delete(id uint) {
	performance := Performance{}
	performance.ID = id

	p.Db.Delete(&performance)
}
