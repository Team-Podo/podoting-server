package repository

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	ID             uint `json:"id" gorm:"primarykey"`
	Name           string
	Birth          *time.Time      `json:"birth"`
	ProfileImage   *File           `json:"profileFile" gorm:"foreignkey:ProfileImageID"`
	ProfileImageID *uint           `json:"-"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	DeletedAt      *gorm.DeletedAt `json:"-" gorm:"index"`
}

func (p *Person) SetBirth(birth string) error {
	t, err := time.Parse("2006-01-02", birth)
	if err != nil {
		return err
	}

	p.Birth = &t

	return nil
}

type PersonRepository struct {
	DB *gorm.DB
}

func (r *PersonRepository) FindAll() ([]Person, error) {
	var persons []Person

	if err := r.DB.Find(&persons).Error; err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *PersonRepository) FindByID(id uint) (*Person, error) {
	var person Person

	if err := r.DB.First(&person, id).Error; err != nil {
		return nil, err
	}

	return &person, nil
}

func (r *PersonRepository) Create(person *Person) error {
	if err := r.DB.Create(person).Error; err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) Update(person *Person) error {
	if err := r.DB.Save(person).Error; err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) Delete(id uint) error {
	if err := r.DB.Delete(&Person{ID: id}).Error; err != nil {
		return err
	}

	return nil
}
