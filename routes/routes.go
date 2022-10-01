package routes

import (
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance/cast"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance/place"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance/place/area"
	"github.com/Team-Podo/podoting-server/endpoints/admin/product"
	"github.com/Team-Podo/podoting-server/endpoints/admin/schedule"
	"github.com/Team-Podo/podoting-server/endpoints/musical"
	"github.com/Team-Podo/podoting-server/endpoints/musical/seat"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Routes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://podoting.com",
			"https://www.podoting.com",
			"https://app.podoting.com",
			"https://partner.podoting.com",
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		MaxAge:       12 * time.Hour,
	}))

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
			products.POST("/:id/main-image", product.UploadMainImage)
		}

		performances := admin.Group("/performances")
		{
			performances.GET("/", performance.Get)
			performances.GET("/:id", performance.Find)
			performances.POST("/", performance.Create)
			performances.POST("/:id/thumbnail", performance.UploadThumbnailImage)
			performances.PUT("/:id", performance.Update)
			performances.DELETE("/", performance.Delete)

			casts := performances.Group("/:id/casts")
			{
				casts.POST("/:cast_id/profile-image", cast.UploadProfileImage)
			}

			places := performances.Group("/:id/places")
			{
				places.GET("/", place.Find)
				places.POST("/:place_id/place-image", place.UploadPlaceImage)
			}
		}

		places := admin.Group("/places")
		{
			places.GET("/:id", place.Find)
			places.GET("/", place.Get)
			places.POST("/", place.Create)
			places.PUT("/:id", place.Update)
			places.DELETE("/:id", place.Delete)

			areas := places.Group("/:id/areas")
			{
				areas.GET("/", area.Get)
				areas.GET("/:area_id", area.Find)
				areas.POST("/", area.Create)
				areas.POST("/:area_id/background-image", area.UploadAreaImage)
				areas.PUT("/:area_id", area.Update)
				areas.DELETE("/:area_id", area.Delete)
			}
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

	musicalGroup := r.Group("/musical")
	{
		musicalGroup.GET("/:id", musical.Find)
		musicalGroup.GET("/:id/schedules/:schedule_uuid/seats", seat.Get)
	}
}
