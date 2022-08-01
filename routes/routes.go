package routes

import (
	"github.com/Team-Podo/podoting-server/endpoints"
	"github.com/Team-Podo/podoting-server/endpoints/admin/area"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance"
	"github.com/Team-Podo/podoting-server/endpoints/admin/place"
	"github.com/Team-Podo/podoting-server/endpoints/admin/product"
	"github.com/Team-Podo/podoting-server/endpoints/admin/schedule"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

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

		performances := admin.Group("/performances")
		{
			performances.GET("/", performance.Get)
			performances.GET("/:id", performance.Find)
			performances.POST("/", performance.Create)
			performances.PUT("/:id", performance.Update)
			performances.DELETE("/", performance.Delete)
		}

		schedules := admin.Group("/schedules")
		{
			schedules.GET("/", schedule.Get)
			schedules.GET("/:uuid", schedule.Find)
			schedules.POST("/", schedule.Create)
			schedules.PUT("/:uuid", schedule.Update)
			schedules.DELETE("/", schedule.Delete)
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
