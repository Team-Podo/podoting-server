package routes

import (
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance/cast"
	"github.com/Team-Podo/podoting-server/endpoints/admin/product"
	"github.com/Team-Podo/podoting-server/endpoints/admin/schedule"
	"github.com/Team-Podo/podoting-server/endpoints/musical"
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
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
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
			performances.PUT("/:id", performance.Update)
			performances.DELETE("/", performance.Delete)

			casts := performances.Group("/:id/casts")
			{
				casts.POST("/:cast_id/profile-image", cast.UploadProfileImage)
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
	}

}
