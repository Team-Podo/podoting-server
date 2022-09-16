package repository

type Point struct {
	ID uint    `gorm:"primarykey"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
}
