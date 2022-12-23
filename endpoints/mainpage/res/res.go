package res

import "github.com/Team-Podo/podoting-server/repository"

type MainPageResponse struct {
	Performances []Performance `json:"performances"`
}

type Performance struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	ThumbUrl    *string     `json:"thumbUrl"`
	RunningTime string      `json:"runningTime"`
	StartDate   string      `json:"startDate"`
	EndDate     string      `json:"endDate"`
	Rating      string      `json:"rating"`
	Place       *Place      `json:"place"`
	Schedules   []Schedule  `json:"schedules"`
	Cast        []Cast      `json:"cast"`
	Contents    []Content   `json:"contents"`
	SeatGrades  []SeatGrade `json:"seatGrades"`
}

type Place struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Schedule struct {
	UUID string         `json:"uuid"`
	Date string         `json:"date"`
	Time string         `json:"time"`
	Cast []ScheduleCast `json:"cast"`
}

type ScheduleCast struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Cast struct {
	ID      uint    `json:"id"`
	Profile Profile `json:"profile"`
	Name    string  `json:"name"`
	Role    string  `json:"role"`
}

type Profile struct {
	Url string `json:"url"`
}

type Content struct {
	Content string `json:"content"`
}

type SeatGrade struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (m MainPageResponse) Of(performances []repository.Performance) MainPageResponse {
	var response MainPageResponse
	response.Performances = make([]Performance, len(performances))
	for i := 0; i < len(performances); i++ {
		p := performances[i]
		response.Performances[i] = Performance{
			ID:          p.ID,
			Title:       p.Title,
			ThumbUrl:    p.GetThumbnailURL(),
			RunningTime: p.RunningTime,
			StartDate:   p.StartDate,
			EndDate:     p.EndDate,
			Rating:      p.Rating,
			Place:       getPlace(p.Place),
			Schedules:   getSchedules(p.Schedules),
			Cast:        getCasts(p.Casts),
			Contents:    getContents(p.Contents),
			SeatGrades:  getSeatGrades(p.SeatGrades),
		}
	}

	return response
}

func getPlace(place *repository.Place) *Place {
	if place == nil {
		return nil
	}

	return &Place{
		ID:    place.ID,
		Name:  place.Name,
		Image: place.Name,
	}
}

func getSchedules(schedules []repository.Schedule) []Schedule {
	s := make([]Schedule, len(schedules))
	for i, schedule := range schedules {
		s[i] = Schedule{
			UUID: schedule.UUID,
			Date: schedule.Date,
			Time: schedule.Time.String,
			Cast: getScheduleCasts(schedule.Casts),
		}
	}
	return s
}

func getScheduleCasts(casts []repository.Cast) []ScheduleCast {
	c := make([]ScheduleCast, len(casts))
	for i, cast := range casts {
		c[i] = ScheduleCast{
			ID:   cast.ID,
			Name: cast.Person.Name,
		}
	}
	return c
}

func getCasts(casts []repository.Cast) []Cast {
	c := make([]Cast, len(casts))
	for i, cast := range casts {
		c[i] = Cast{
			ID: cast.ID,
			Profile: Profile{
				Url: cast.ProfileImageURL(),
			},
			Name: cast.Person.Name,
			Role: cast.Character.Name,
		}
	}
	return c
}

func getContents(contents []repository.PerformanceContent) []Content {
	c := make([]Content, len(contents))

	for i, content := range contents {
		c[i] = Content{
			Content: content.Content,
		}
	}

	return c
}

func getSeatGrades(seatGrades []repository.SeatGrade) []SeatGrade {
	s := make([]SeatGrade, len(seatGrades))

	for i, grade := range seatGrades {
		s[i] = SeatGrade{
			ID:    grade.ID,
			Name:  grade.Name,
			Price: grade.Price,
		}
	}

	return s
}
