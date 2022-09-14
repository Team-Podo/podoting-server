package routes

import (
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance/cast"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance/place"
	"github.com/Team-Podo/podoting-server/endpoints/admin/product"
	"github.com/Team-Podo/podoting-server/endpoints/admin/schedule"
	"github.com/Team-Podo/podoting-server/endpoints/musical"
	"github.com/Team-Podo/podoting-server/utils"
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
			performances.PUT("/:id", performance.Update)
			performances.DELETE("/", performance.Delete)

			casts := performances.Group("/:id/casts")
			{
				casts.POST("/:cast_id/profile-image", cast.UploadProfileImage)
			}

			places := performances.Group("/:id/places")
			{
				places.POST("/:place_id/place-image", place.UploadPlaceImage)
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
		musicalGroup.GET("/:id/schedules/:scheduleUUID/seats", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"backgroundImage": "https://podoting.com/images/seat-background.png",
				"seats": []utils.MapSlice{
					utils.BuildMapSliceByMap(map[string]any{
						"uuid": "a6c6a191-7ef9-4695-b89b-d76bf2e31d08",
						"point": utils.BuildMapSliceByMap(map[string]any{
							"x": 530.2091,
							"y": 620.1124,
						}),
						"name": "I 열 11번 좌석",
						"grade": utils.BuildMapSliceByMap(map[string]any{
							"id":   102,
							"name": "VIP",
						}),
						"price": 99000,
						"color": "#7748F4",
					}),
					utils.BuildMapSliceByMap(map[string]any{
						"uuid": "a6c6a191-7ef9-4695-b89b-d76bf2e31d09",
						"point": utils.BuildMapSliceByMap(map[string]any{
							"x": 534.2091,
							"y": 620.1124,
						}),
						"name": "I 열 12번 좌석",
						"grade": utils.BuildMapSliceByMap(map[string]any{
							"id":   102,
							"name": "VIP",
						}),
						"price": 99000,
						"color": "#7748F4",
					}),
					utils.BuildMapSliceByMap(map[string]any{
						"uuid": "a6c6a191-7ef9-4695-b89b-d76bf2e31d10",
						"point": utils.BuildMapSliceByMap(map[string]any{
							"x": 538.2091,
							"y": 620.1124,
						}),
						"name": "I 열 13번 좌석",
						"grade": utils.BuildMapSliceByMap(map[string]any{
							"id":   102,
							"name": "VIP",
						}),
						"price": 99000,
						"color": "#7748F4",
					}),
				},
			})
		})
	}
}
