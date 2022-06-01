package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/endpoints"
)

func Routes(r *gin.Engine) {
	products := r.Group("/products")
	{
		products.GET("/:id", endpoints.GetProduct)
		products.POST("/", endpoints.SaveProduct)
		products.DELETE("/:id", endpoints.DeleteProduct)
	}

	places := r.Group("/places")
	{
		places.GET("/", endpoints.GetPlaces)
		places.GET("/:placeId", endpoints.GetPlace)
		places.POST("/", endpoints.SavePlace)
		places.PUT("/:placeId", endpoints.UpdatePlace)

		areas := r.Group("/:placeId/areas")
		{
			areas.GET("/:areaId", endpoints.GetArea)
		}
	}

}
