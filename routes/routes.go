package routes

import (
	"github.com/Team-Podo/podoting-server/endpoints/admin/cast"
	"github.com/Team-Podo/podoting-server/endpoints/admin/character"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance"
	performanceCast "github.com/Team-Podo/podoting-server/endpoints/admin/performance/cast"
	performanceCharacter "github.com/Team-Podo/podoting-server/endpoints/admin/performance/character"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance/content"
	"github.com/Team-Podo/podoting-server/endpoints/admin/performance/grades"
	performanceSchedule "github.com/Team-Podo/podoting-server/endpoints/admin/performance/schedule"
	adminSeats "github.com/Team-Podo/podoting-server/endpoints/admin/performance/seat"
	"github.com/Team-Podo/podoting-server/endpoints/admin/person"
	"github.com/Team-Podo/podoting-server/endpoints/admin/place"
	"github.com/Team-Podo/podoting-server/endpoints/admin/place/area"
	"github.com/Team-Podo/podoting-server/endpoints/admin/product"
	"github.com/Team-Podo/podoting-server/endpoints/book"
	"github.com/Team-Podo/podoting-server/endpoints/mainpage"
	"github.com/Team-Podo/podoting-server/endpoints/musical"
	"github.com/Team-Podo/podoting-server/endpoints/musical/seat"
	"github.com/Team-Podo/podoting-server/endpoints/mypage"
	"github.com/Team-Podo/podoting-server/endpoints/order"
	"github.com/Team-Podo/podoting-server/middleware"
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
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
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

		casts := admin.Group("/casts")
		{
			casts.GET("/:id", cast.FindOne)
			casts.DELETE("/:id", cast.Delete)
			casts.PUT("/:id", cast.Update)
			casts.POST("/:id/profile-image", cast.UploadProfileImage)
		}

		characters := admin.Group("/characters")
		{
			characters.PUT("/:id", character.Update)
			characters.DELETE("/:id", character.Delete)
		}

		people := admin.Group("/people")
		{
			people.GET("/", person.Get)
			people.GET("/:id", person.Find)
			people.POST("/", person.Create)
			people.PUT("/:id", person.Update)
			people.DELETE("/:id", person.Delete)
		}

		schedule := admin.Group("/schedules")
		{
			schedule.PUT("/:uuid", performanceSchedule.Update)
			schedule.DELETE("/:uuid", performanceSchedule.Delete)
		}

		performances := admin.Group("/performances")
		{
			performances.GET("/", performance.Get)
			performances.GET("/:id", performance.Find)
			performances.POST("/", performance.Create)
			performances.POST("/:id/thumbnail", performance.UploadThumbnailImage)
			performances.PUT("/:id", performance.Update)
			performances.DELETE("/:id", performance.Delete)

			places := performances.Group("/:id/places")
			{
				places.GET("/", place.Find)
				places.POST("/:place_id/place-image", place.UploadPlaceImage)
			}

			performanceCasts := performances.Group("/:id/casts")
			{
				performanceCasts.GET("/", performanceCast.Get)
				performanceCasts.POST("/many", performanceCast.CreateMany)
			}

			performanceCharacters := performances.Group("/:id/characters")
			{
				performanceCharacters.GET("/", performanceCharacter.Get)
				performanceCharacters.POST("/", performanceCharacter.Create)
			}

			performanceContents := performances.Group("/:id/contents")
			{
				performanceContents.GET("/", content.Find)
				performanceContents.GET("/:content_id", content.FindOne)
				performanceContents.POST("/", content.Create)
				performanceContents.POST("/image", content.UploadContentImage)
				performanceContents.PUT("/:content_id", content.Update)
				performanceContents.DELETE("/:content_id", content.Delete)
			}

			performanceSchedules := performances.Group("/:id/schedules")
			{
				performanceSchedules.GET("/", performanceSchedule.Get)
				performanceSchedules.POST("/", performanceSchedule.Create)
			}

			performanceAreas := performances.Group("/:id/areas")
			{
				performanceAreas.POST("/:area_id/match", performance.MatchArea)

				seats := performanceAreas.Group("/:area_id/seats")
				{
					seats.GET("/", adminSeats.Get)
					seats.POST("/", adminSeats.Save)
				}
			}

			performanceSeatGrades := performances.Group("/:id/seat-grades")
			{
				performanceSeatGrades.GET("/", grades.GetByPerformanceID)
				performanceSeatGrades.POST("/", grades.Create)
				performanceSeatGrades.PUT("/:seat_grade_id", grades.Update)
				performanceSeatGrades.DELETE("/:seat_grade_id", grades.Delete)
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
	}

	musicalGroup := r.Group("/musical")
	{
		musicalGroup.GET("/:id", musical.Find)
		musicalGroup.GET("/:id/schedules/:schedule_uuid/seats", seat.Get)
	}

	bookGroup := r.Group("/book", middleware.AuthMiddleware)
	{
		bookGroup.POST("/:schedule_uuid", book.Book)
	}

	mypageGroup := r.Group("/mypage", middleware.AuthMiddleware)
	{
		mypageGroup.GET("/order-history", mypage.GetOrderHistory)
	}

	orderGroup := r.Group("/order", middleware.AuthMiddleware)
	{
		orderGroup.PATCH("/cancel/:id", order.CancelOrder)
		orderGroup.PATCH("/cancel/:id/details/:order_detail_id", order.CancelOrderDetail)
	}

	mainpageGroup := r.Group("/mainpage")
	{
		mainpageGroup.GET("/", mainpage.Index)
	}
}
