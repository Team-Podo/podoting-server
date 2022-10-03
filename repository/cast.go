package repository

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"time"
)

type Cast struct {
	ID             uint       `json:"id" gorm:"primarykey"`
	Character      *Character `json:"character" gorm:"foreignkey:CharacterID"`
	CharacterID    uint       `json:"-"`
	Person         *Person    `json:"person" gorm:"foreignkey:PersonID"`
	PersonID       uint       `json:"-"`
	ProfileImage   *File      `json:"profileImage" gorm:"foreignkey:ProfileImageID"`
	ProfileImageID *uint      `json:"-"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"-" gorm:"index"`
}

func (c *Cast) ProfileImageURL() string {
	if c.ProfileImage == nil {
		return ""
	}

	return os.Getenv("CDN_URL") + "/" + c.ProfileImage.Path
}

type CastRepository struct {
	DB *gorm.DB
}

func (c *CastRepository) Get() ([]Cast, error) {
	var casts []Cast

	err := c.DB.Find(&casts).Error

	if err != nil {
		return nil, err
	}

	return casts, nil
}

func (c *CastRepository) GetByPerformanceID(performanceID uint) ([]Cast, error) {
	var performance Performance
	performance.ID = performanceID

	err := c.DB.
		Model(&performance).
		Joins("Character").
		Joins("Person").
		Joins("ProfileImage").
		Association("Casts").
		Find(&performance.Casts)

	fmt.Println("character", performance.Casts[0].Person)

	if err != nil {
		return nil, err
	}

	return performance.Casts, nil
}

func (c *CastRepository) FindByID(id uint) (*Cast, error) {
	var cast Cast
	cast.ID = id

	err := c.DB.
		Joins("Character").
		Joins("Person").
		Joins("ProfileImage").
		First(&cast).
		Error

	if err != nil {
		return nil, err
	}

	return &cast, nil
}

func (c *CastRepository) Create(cast *Cast) error {
	if cast.ProfileImage != nil {
		cast.ProfileImageID = &cast.ProfileImage.ID
	}

	err := c.DB.Create(cast).Error

	if err != nil {
		return err
	}

	return nil
}

func (c *CastRepository) CreateMany(casts []Cast) error {
	err := c.DB.Save(casts).Error

	if err != nil {
		return err
	}

	return nil
}

func (c *CastRepository) LinkPerformances(performanceCasts []PerformanceCast) error {
	err := c.DB.Save(performanceCasts).Error

	if err != nil {
		return err
	}

	return nil
}

func (c *CastRepository) Delete(id uint) error {
	err := c.DB.Delete(&Cast{ID: id}).Error

	if err != nil {
		return err
	}

	return nil
}

func (c *CastRepository) GetCastsByPerformanceID(id uint) ([]*Cast, error) {
	var casts []*Cast

	return casts, nil
}

func (c *CastRepository) Update(cast *Cast) error {
	if cast.ProfileImage != nil {
		cast.ProfileImageID = &cast.ProfileImage.ID
	}

	err := c.DB.Model(&Cast{ID: cast.ID}).Updates(cast).Error

	if err != nil {
		return err
	}

	return nil
}
