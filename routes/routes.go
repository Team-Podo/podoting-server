package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/endpoints/product"
)

func Routes(r *gin.Engine) {
	r.GET("/products/:id", product.GetProduct)
}
