package musical

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/musical"
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

	placeCH := make(chan *musical.Place)
	go getPlace(performance, placeCH)

	scheduleCH := make(chan []musical.Schedule)
	go getSchedules(performance.ID, scheduleCH)

	castCH := make(chan []musical.Cast)
	go getCasts(performance.ID, castCH)

	contentCH := make(chan []musical.Content)
	go getContents(performance.ID, contentCH)

	seatGradeCH := make(chan []musical.SeatGrade)
	go getSeatGrades(performance.ID, seatGradeCH)

	c.JSON(200, musical.Musical{
		ID:          performance.ID,
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
		SeatGrades:  <-seatGradeCH,
	})
}

func getCasts(id uint, ch chan []musical.Cast) {
	casts := repositories.performance.GetCastsByID(id)
	var result []musical.Cast

	for _, cast := range casts {
		result = append(result, musical.Cast{
			ID: cast.ID,
			Profile: musical.Profile{
				Url: cast.ProfileImageURL(),
			},
			Name: cast.Person.Name,
			Role: cast.Character.Name,
		})
	}

	ch <- result
}

func getPlace(p *repository.Performance, ch chan *musical.Place) {
	if p.Place == nil {
		ch <- nil
		return
	}

	ch <- &musical.Place{
		ID:    p.Place.ID,
		Name:  p.Place.Name,
		Image: p.Place.ImageURL(),
	}
}

func getSchedules(id uint, ch chan []musical.Schedule) {
	schedules := repositories.performance.GetSchedulesByID(id)
	if schedules == nil {
		ch <- nil
		return
	}

	var result []musical.Schedule

	for _, schedule := range schedules {
		var casts []musical.ScheduleCast

		for _, cast := range schedule.Casts {
			casts = append(casts, musical.ScheduleCast{
				ID:   cast.ID,
				Name: cast.Person.Name,
			})
		}

		result = append(result, musical.Schedule{
			UUID: schedule.UUID,
			Date: schedule.Date,
			Time: schedule.Time.String,
			Cast: casts,
		})
	}

	ch <- result
}

func getContents(id uint, ch chan []musical.Content) {
	contents := repositories.performance.GetContentsByID(id)
	if contents == nil {
		ch <- nil
		return
	}

	var result []musical.Content

	for _, content := range contents {
		result = append(result, musical.Content{
			UUID:    content.UUID,
			Title:   content.Title,
			Content: content.Content,
		})
	}

	ch <- result
}

func getSeatGrades(id uint, ch chan []musical.SeatGrade) {
	seatGrades := repositories.performance.GetSeatGradesByID(id)
	if seatGrades == nil {
		ch <- nil
		return
	}

	var result []musical.SeatGrade

	for _, seatGrade := range seatGrades {
		result = append(result, musical.SeatGrade{
			ID:    seatGrade.ID,
			Name:  seatGrade.Name,
			Price: seatGrade.Price,
		})
	}

	ch <- result
}
