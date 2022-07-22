package models

type Product interface {
	GetId() uint
	GetTitle() string
	GetPlace() Place
	IsNil() bool
	IsNotNil() bool
	GetCreatedAt() string
	GetUpdatedAt() string
}

type ProductRepository interface {
	Get(query map[string]any) []Product
	Find(id uint) Product
	Save(product Product) Product
	Update(product Product) Product
	Delete(id uint)
	GetTotal(query map[string]any) int64
}
