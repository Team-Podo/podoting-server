package repository

type Place struct {
	Model
	Title    string
	Products *[]Product `gorm:"foreignkey:PlaceId"`
	Areas    *[]Area    `gorm:"foreignkey:PlaceId"`
}
