package models

type Product interface {
	GetId() uint
	GetTitle() string
}

type ProductRepository interface {
	Get() []Product
	GetProductById(id uint) Product
	SaveProduct(product Product) Product
	DeleteProductById(id uint)
}
