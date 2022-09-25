package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/utils"
	"strconv"
	"time"
)

type Performance struct {
	ID          uint                  `json:"id" gorm:"primarykey"`
	Product     *Product              `json:"product" gorm:"foreignkey:ProductID"`
	ProductID   *uint                 `json:"-"`
	Thumbnail   *File                 `json:"thumbnail" gorm:"foreignkey:ThumbnailID"`
	ThumbnailID *uint                 `json:"-"`
	Place       *Place                `json:"place" gorm:"foreignkey:PlaceID"`
	PlaceID     *uint                 `json:"-"`
	Areas       []*Area               `json:"areas" gorm:"many2many:performance_areas;"`
	MainArea    *Area                 `json:"main_area" gorm:"foreignkey:MainAreaID"`
	MainAreaID  *uint                 `json:"main_area_id"`
	Casts       []*Cast               `gorm:"many2many:performance_casts;"`
	Schedules   []*Schedule           `gorm:"foreignkey:PerformanceID"`
	Contents    []*PerformanceContent `gorm:"foreignkey:PerformanceID"`
	Title       string                `json:"title"`
	RunningTime string                `json:"runningTime"`
	StartDate   string                `json:"startDate"`
	EndDate     string                `json:"endDate"`
	Rating      string                `json:"rating"`
	CreatedAt   time.Time             `json:"createdAt"`
	UpdatedAt   time.Time             `json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt       `json:"-" gorm:"index"`
}

func (p *Performance) GetFileURL() string {
	if p.Product.IsNil() {
		return ""
	}

	return p.Product.GetFileURL()
}

type PerformanceRepository struct {
	DB *gorm.DB
}

func (p *PerformanceRepository) GetWithQueryMap(query map[string]any) []Performance {
	var performances []Performance

	limit, _ := strconv.Atoi(utils.ToString(query["limit"]))
	offset, _ := strconv.Atoi(utils.ToString(query["offset"]))
	reversed, _ := query["reversed"].(bool)

	result := p.DB.
		Debug().
		Limit(limit).
		Offset(offset).
		Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: reversed}).
		Find(&performances)

	if result.Error != nil {
		return nil
	}

	if len(performances) == 0 {
		return nil
	}

	return performances
}

func (p *PerformanceRepository) GetTotalWithQueryMap(query map[string]any) int64 {
	var count int64

	limit, _ := strconv.Atoi(utils.ToString(query["limit"]))
	offset, _ := strconv.Atoi(utils.ToString(query["offset"]))
	reversed, _ := query["reversed"].(bool)

	p.DB.Limit(limit).
		Offset(offset).
		Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: reversed}).
		Model(&Performance{}).Count(&count)

	return count
}

func (p *PerformanceRepository) applyAllQuery(query map[string]any) {
	p.applyReversedQuery(query)
	p.applyLimitQuery(query)
	p.applyOffsetQuery(query)
}

func (p *PerformanceRepository) applyReversedQuery(query map[string]any) {
	if query["reversed"] == true {
		fmt.Println("reversed")
		p.DB = p.DB.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: true})
	} else {
		fmt.Println("not reversed")
		p.DB = p.DB.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: false})
	}
}

func (p *PerformanceRepository) applyLimitQuery(query map[string]any) {
	if query["limit"] != nil {
		limit, _ := strconv.Atoi(utils.ToString(query["limit"]))
		p.DB = p.DB.Limit(limit)
	} else {
		p.DB = p.DB.Limit(10)
	}
}

func (p *PerformanceRepository) applyOffsetQuery(query map[string]any) {
	if query["offset"] != nil {
		offset, _ := strconv.Atoi(utils.ToString(query["offset"]))
		p.DB = p.DB.Offset(offset)
	} else {
		p.DB = p.DB.Offset(0)
	}
}

func (p *PerformanceRepository) FindByID(id uint) *Performance {
	performance := Performance{
		ID: id,
	}

	err := p.DB.
		Preload("Place.PlaceImage").
		Preload("Product.File").
		Joins("Thumbnail").
		First(&performance).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		return nil
	}

	return &performance
}

func (p *PerformanceRepository) CheckMainAreaExistsByID(performanceID uint) (uint, error) {
	var performance Performance

	p.DB.Where("main_area_id is not null").Find(&performance, performanceID)

	if performance.ID == 0 {
		return 0, errors.New("main area not exists")
	}

	return *performance.MainAreaID, nil
}

func (p *PerformanceRepository) Save(performance *Performance) error {
	err := p.DB.Create(performance).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *PerformanceRepository) Update(performance *Performance) error {
	err := p.DB.First(&Performance{}, performance.ID).Error

	if err != nil {
		return err
	}

	p.DB.Model(&Performance{ID: performance.ID}).Updates(performance)

	return nil
}

func (p *PerformanceRepository) Delete(id uint) error {
	performance := Performance{}
	performance.ID = id

	err := p.DB.Delete(&performance).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *PerformanceRepository) GetCastsByID(id uint) []*Cast {
	var performance Performance
	performance.ID = id

	err := p.DB.
		Model(&performance).
		Joins("Character").
		Joins("Person").
		Joins("ProfileImage").
		Association("Casts").
		Find(&performance.Casts)

	if err != nil {
		return nil
	}

	return performance.Casts
}

func (p *PerformanceRepository) GetSchedulesByID(id uint) []*Schedule {
	var performance Performance
	performance.ID = id

	err := p.DB.
		Model(&performance).
		Order("date asc").
		Order("time asc").
		Preload("Casts.Person").
		Association("Schedules").
		Find(&performance.Schedules)

	if err != nil {
		return nil
	}

	return performance.Schedules
}

func (p *PerformanceRepository) GetContentsByID(id uint) []*PerformanceContent {
	var performance Performance
	performance.ID = id

	err := p.DB.
		Model(&performance).
		Order("priority asc").
		Association("Contents").
		Find(&performance.Contents)

	if err != nil {
		return nil
	}

	return performance.Contents
}
