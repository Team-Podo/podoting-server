package repository

import (
	"gorm.io/gorm"
	"os"
	"time"
)

type Place struct {
	ID           uint            `json:"id" gorm:"primarykey"`
	Name         string          `json:"name"`
	Location     *Location       `json:"location"`
	LocationID   *uint           `json:"-"`
	PlaceImage   *File           `json:"placeImage" gorm:"foreignkey:PlaceImageID"`
	PlaceImageID *uint           `json:"-"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	DeletedAt    *gorm.DeletedAt `json:"-" gorm:"index"`
}

func (p *Place) ImageURL() string {
	if p.PlaceImage.IsNil() {
		return ""
	}

	return os.Getenv("CDN_URL") + "/" + p.PlaceImage.Path
}

type PlaceRepository struct {
	DB *gorm.DB
}

func (r *PlaceRepository) Create(place *Place) error {
	if err := r.DB.Debug().Create(place).Error; err != nil {
		return err
	}

	return nil
}

func (r *PlaceRepository) FindByID(id uint) (*Place, error) {
	var place Place
	if err := r.DB.Joins("Location").First(&place, id).Error; err != nil {
		return nil, err
	}

	return &place, nil
}

func (r *PlaceRepository) FindAll() ([]Place, error) {
	var places []Place
	if err := r.DB.Joins("Location").Find(&places).Error; err != nil {
		return nil, err
	}

	return places, nil
}

func (r *PlaceRepository) Update(place *Place) error {
	if err := r.DB.Model(Place{ID: place.ID}).Updates(Place{
		Name: place.Name,
	}).Error; err != nil {
		return err
	}

	if place.Location != nil {
		if err := r.DB.Save(place.Location).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *PlaceRepository) Delete(id uint) error {
	if err := r.DB.Delete(&Place{}, id).Error; err != nil {
		return err
	}

	return nil
}
