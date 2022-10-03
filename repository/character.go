package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Character struct {
	ID            uint         `json:"id" gorm:"primarykey"`
	Performance   *Performance `json:"performance" gorm:"foreignkey:PerformanceID"`
	PerformanceID uint         `json:"-"`
	Name          string
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	DeletedAt     *gorm.DeletedAt `json:"-" gorm:"index"`
}

type CharacterRepository struct {
	DB *gorm.DB
}

func (r *CharacterRepository) Create(character *Character) error {
	if err := r.DB.Debug().Create(character).Error; err != nil {
		return err
	}

	return nil
}

func (r *CharacterRepository) FindByPerformanceID(performanceID uint) ([]Character, error) {
	var characters []Character

	if err := r.DB.
		Where("performance_id = ?", performanceID).
		Find(&characters).Error; err != nil {
		return nil, err
	}

	return characters, nil
}

func (r *CharacterRepository) Update(character *Character) error {
	if err := r.DB.Model(Character{ID: character.ID}).Updates(Character{
		Name: character.Name,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *CharacterRepository) Delete(id uint) error {
	if err := r.DB.Model(Cast{
		CharacterID: id,
	}).First(&Cast{}).Error; err == nil {
		return errors.Wrap(err, "character is used by cast")
	}

	if err := r.DB.Delete(&Character{ID: id}).Error; err != nil {
		return errors.Wrap(err, "database error: failed to delete character")
	}

	return nil
}
