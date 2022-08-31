package musical

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var repositories Repository

type Repository struct {
	performance models.PerformanceRepository
}

func init() {
	repositories = Repository{
		performance: &repository.PerformanceRepository{DB: database.Gorm},
	}
}

func Find(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be Integer")
		return
	}

	performance := repositories.performance.FindByID(uint(intId))
	if performance == nil {
		c.JSON(http.StatusNotFound, "Musical Not Found")
		return
	}

	placeCH := make(chan *models.MusicalPlace)
	go getPlace(performance, placeCH)

	scheduleCH := make(chan []models.MusicalSchedule)
	go getSchedules(performance.ID, scheduleCH)

	castCH := make(chan []models.Cast)
	go getCasts(performance.ID, castCH)

	contentCH := make(chan []models.MusicalContent)
	go getContents(performance.ID, contentCH)

	musical := models.Musical{
		Id:          performance.ID,
		Title:       performance.Title,
		ThumbUrl:    performance.GetFileURL(),
		RunningTime: performance.RunningTime,
		StartDate:   performance.StartDate,
		EndDate:     performance.EndDate,
		Rating:      performance.Rating,
		Place:       <-placeCH,
		Schedules:   <-scheduleCH,
		Cast:        <-castCH,
		Contents:    <-contentCH,
	}

	c.JSON(200, musical)
}

func getCasts(id uint, ch chan []models.Cast) {
	casts := repositories.performance.GetCastsByID(id)
	var result []models.Cast

	for _, cast := range casts {
		result = append(result, models.Cast{
			Id: cast.ID,
			Profile: models.Profile{
				Url: cast.ProfileImageURL(),
			},
			Name: cast.Person.Name,
			Role: cast.Character.Name,
		})
	}

	ch <- result
}

func getPlace(p *repository.Performance, ch chan *models.MusicalPlace) {
	if p.Place == nil {
		ch <- nil
		return
	}

	ch <- &models.MusicalPlace{
		ID:    p.Place.ID,
		Name:  p.Place.Name,
		Image: p.Place.ImageURL(),
	}
}

func getSchedules(id uint, ch chan []models.MusicalSchedule) {
	schedules := repositories.performance.GetSchedulesByID(id)
	if schedules == nil {
		ch <- nil
		return
	}

	var result []models.MusicalSchedule

	for _, schedule := range schedules {
		var casts []models.MusicalScheduleCast

		for _, cast := range schedule.Casts {
			casts = append(casts, models.MusicalScheduleCast{
				ID:   cast.ID,
				Name: cast.Person.Name,
			})
		}

		result = append(result, models.MusicalSchedule{
			UUID: schedule.UUID,
			Date: schedule.Date,
			Time: schedule.Time.String,
			Cast: casts,
		})
	}

	ch <- result
}

func getContents(id uint, ch chan []models.MusicalContent) {
	contents := repositories.performance.GetContentsByID(id)
	if contents == nil {
		ch <- nil
		return
	}

	var result []models.MusicalContent

	for _, content := range contents {
		result = append(result, models.MusicalContent{
			Uuid:    content.UUID,
			Title:   content.Title,
			Content: content.Content,
		})
	}

	ch <- result
}
