package models

type Product interface {
	GetId() uint
	GetTitle() string
	GetPlace() Place
}

type ProductRepository interface {
	Get() []Product
	Find(id uint) Product
	Save(product Product) Product
	Update(product Product) Product
	Delete(id uint)
}
