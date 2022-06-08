package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/endpoints"
	"github.com/kwanok/podonine/endpoints/admin/area"
	"github.com/kwanok/podonine/endpoints/admin/place"
	"github.com/kwanok/podonine/endpoints/admin/product"
)

func Routes(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		products := admin.Group("/products")
		{
			products.GET("/", product.Get)
			products.GET("/:id", product.Find)
			products.POST("/", product.Create)
		}

		places := admin.Group("/places")
		{
			places.GET("/", place.Get)
			places.GET("/:id", place.Find)
			places.POST("/", place.Create)
		}

		areas := admin.Group("/areas")
		{
			areas.GET("/", area.Get)
			areas.GET("/:id", area.Find)
			areas.POST("/", area.Create)
		}
	}

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
