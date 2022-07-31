package repository

import (
	"errors"
	"github.com/Team-Podo/podoting-server/models"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"log"
	"strconv"
	"time"
)

type Performance struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	ProductID uint            `json:"-"`
	Product   *Product        `json:"product" gorm:"foreignkey:ProductID"`
	Title     string          `json:"title"`
	StartDate string          `json:"startDate"`
	EndDate   string          `json:"endDate"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
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

func (p *Performance) GetCreatedAt() string {
	return p.CreatedAt.Format("2006-01-02 15:04:05")
}

func (p *Performance) GetUpdatedAt() string {
	return p.UpdatedAt.Format("2006-01-02 15:04:05")
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

	err := db.Find(&performances).Error

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

func (p *PerformanceRepository) Save(performanceModel models.Performance) models.Performance {
	performance := Performance{
		ProductID: performanceModel.GetProduct().GetId(),
		Title:     performanceModel.GetTitle(),
		StartDate: performanceModel.GetStartDate(),
		EndDate:   performanceModel.GetEndDate(),
	}

	result := p.Db.Create(&performance)

	if result.Error != nil {
		return nil
	}

	return &performance
}

func (p *PerformanceRepository) Update(performanceModel models.Performance) models.Performance {
	err := p.Db.First(&Performance{}, performanceModel.GetId()).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	p.Db.Model(&Performance{ID: performanceModel.GetId()}).Updates(Performance{
		ProductID: performanceModel.GetProduct().GetId(),
		Title:     performanceModel.GetTitle(),
		StartDate: performanceModel.GetStartDate(),
		EndDate:   performanceModel.GetEndDate(),
	})

	return performanceModel
}

func (p *PerformanceRepository) Delete(id uint) {
	performance := Performance{}
	performance.ID = id

	p.Db.Delete(&performance)
}

func (p *PerformanceRepository) GetTotal(query map[string]any) int64 {

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

	var count int64
	db.Model(&Performance{}).Count(&count)

	return count
}
