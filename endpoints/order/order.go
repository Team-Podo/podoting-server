package order

import (
	"github.com/gin-gonic/gin"
)

type Order struct {
	Product Product `json:"product"`
	Buyer   Buyer   `json:"buyer"`
}

type Product struct {
	Title string `json:"title"`
	Areas []Area `json:"areas"`
}

type Area struct {
	Title string `json:"title"`
	Seats []Seat `json:"seats"`
}

type Seat struct {
	Title string `json:"title"`
}

type Buyer struct {
	ID uint `json:"id"`
}

func create(c *gin.Context) {

}
